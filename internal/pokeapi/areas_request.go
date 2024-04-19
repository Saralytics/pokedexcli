package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageUrl *string) (Location, error) {
	endpoint := "/location/?offset=0&limit=20"
	fullURL := baseURL + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	// Check if cache exists
	fmt.Println(fullURL)
	body, ok := c.cache.Get(fullURL)
	fmt.Println(ok)
	if ok {
		fmt.Println("cache hit")
		locations := Location{}
		err := json.Unmarshal(body, &locations)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return Location{}, err
		}

		return locations, nil
	}

	fmt.Println("cache miss")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	if resp.StatusCode > 399 {
		return Location{}, errors.New("bad status code")
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return Location{}, err
	}

	locations := Location{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return Location{}, err
	}

	c.cache.Add(fullURL, body)
	fmt.Printf("cache added for %s \n", fullURL)
	_, ok = c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache is found")
	}

	return locations, nil
}
