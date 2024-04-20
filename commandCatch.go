package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s ... \n", name)
	resp, err := config.pokeapiclient.CatchPokemon(name)

	if err != nil {
		log.Fatal(err)
	}

	const threshold = 50
	randNum := rand.Intn(resp.BaseExperience)
	if randNum > threshold {
		fmt.Printf("%s escaped! \n", name)
		return nil
	}

	fmt.Printf("%s was caught! \n", name)
	config.caughtPokemon[name] = resp

	return nil

}
