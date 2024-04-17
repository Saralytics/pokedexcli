package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func callbackMap(config *Config) error {
	// get the list of 20 locations
	var url string
	fmt.Println("next: ", config.Next)
	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location/"
	} else {
		url = config.Next
	}

	// base_url := "https://pokeapi.co/api/v2/location/"
	// page := "?offset=20&limit=20"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()
	// if res.StatusCode > 299 {
	// 	fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// 	return errors.New("content no longer available")
	// }
	// Read the response body into a byte slice
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// fmt.Println(body)

	// Unmarshal the JSON data into a struct
	locations := Location{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	fmt.Println("did we set config?: ", config.Next)

	for _, r := range locations.Results {
		fmt.Println("location name is: ", r.Name)
	}
	return nil
}
