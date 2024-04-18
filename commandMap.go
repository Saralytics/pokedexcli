package main

import (
	"fmt"
	"log"

	"github.com/Saralytics/pokedexcli/internal/pokeapi"
)

func callbackMap(config *Config) error {

	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas(config.Next)
	if err != nil {
		log.Fatal(err)
	}

	config.Next = resp.Next
	config.Previous = resp.Previous

	for _, r := range resp.Results {
		fmt.Println("location name is: ", r.Name)
	}
	return nil
}
