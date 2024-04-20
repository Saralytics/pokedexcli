package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (Locations, error) {
	endpoint := "/location-area/?offset=0&limit=20"
	fullURL := baseURL + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	// Check if cache exists
	body, ok := c.cache.Get(fullURL)
	if ok {
		locations := Locations{}
		err := json.Unmarshal(body, &locations)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return Locations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}

	if resp.StatusCode > 399 {
		return Locations{}, errors.New("bad status code")
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return Locations{}, err
	}

	c.cache.Add(fullURL, body)

	return locations, nil
}
