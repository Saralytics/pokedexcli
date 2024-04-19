package main

import (
	"time"

	"github.com/Saralytics/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiclient pokeapi.Client
	Next          *string
	Previous      *string
}

func main() {

	cfg := config{
		pokeapiclient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)

}
