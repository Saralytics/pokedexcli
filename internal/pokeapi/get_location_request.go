package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationName string) (LocationDetails, error) {
	fullURL := baseURL + "/location-area/" + locationName

	// Check if cache exists
	body, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("cache hit")
		locations := LocationDetails{}
		err := json.Unmarshal(body, &locations)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return LocationDetails{}, err
		}

		return locations, nil
	}

	fmt.Println("cache miss")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}

	if resp.StatusCode > 399 {
		return LocationDetails{}, errors.New("bad status code")
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return LocationDetails{}, err
	}

	locations := LocationDetails{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return LocationDetails{}, err
	}

	c.cache.Add(fullURL, body)

	return locations, nil
}
