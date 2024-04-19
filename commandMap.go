package main

import (
	"fmt"
	"log"
)

func callbackMap(config *config) error {

	resp, err := config.pokeapiclient.ListLocationAreas(config.Next)
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

func callbackMapB(config *config) error {

	if config.Previous == nil {
		fmt.Println("You are already on the first page")
		return nil
	}

	resp, err := config.pokeapiclient.ListLocationAreas(config.Previous)
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
