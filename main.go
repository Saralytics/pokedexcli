package main

import (
	"time"

	"github.com/Saralytics/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiclient pokeapi.Client
	Next          *string
	Previous      *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiclient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
