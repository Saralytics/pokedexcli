package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackExplore(config *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	resp, err := config.pokeapiclient.GetLocationArea(name)

	if err != nil {
		log.Fatal(err)
	}
	for _, p := range resp.PokemonEncounters {
		fmt.Println(p.Pokemon.Name)
	}
	return nil
}
