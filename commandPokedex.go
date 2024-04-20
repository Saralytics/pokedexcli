package main

import "fmt"

func callbackPokedex(config *config, args ...string) error {

	fmt.Println("Your Pokedex: ")
	for k := range config.caughtPokemon {
		fmt.Println(k)
	}

	return nil

}
