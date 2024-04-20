package main

import (
	"errors"
	"fmt"
)

func callbackInspect(config *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]
	found_pokemon, ok := config.caughtPokemon[name]
	if !ok {
		fmt.Println("You haven't caught this pokemon yet. Or check if you have a typo.")
		return nil
	}
	fmt.Printf("Name: %s \n", name)
	fmt.Printf("Height: %d \n", found_pokemon.Height)
	fmt.Printf("Weight: %d \n", found_pokemon.Weight)
	fmt.Println("Stats")
	for _, stat := range found_pokemon.Stats {
		fmt.Println("- ", stat.Stat.Name, ":", stat.BaseStat)
	}

	fmt.Println("Types")
	for _, t := range found_pokemon.Types {
		fmt.Println("- ", t.Type.Name)
	}

	return nil
}
