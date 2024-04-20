package main

import "strings"

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	words := strings.Fields(lower)

	return words
}

func getCommands() map[string]clicommand {

	commandMap := map[string]clicommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "display the next 20 locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "display the previous 20 locations",
			callback:    callbackMapB,
		},

		"explore": {
			name:        "explore <location_name>",
			description: "explore the area and see all pokemons there",
			callback:    callbackExplore,
		},

		"catch": {
			name:        "catch <pokemon_name>",
			description: "catch the pokemon",
			callback:    callbackCatch,
		},

		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "inspect the pokemon you have caught",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "see all the pokemons you have",
			callback:    callbackPokedex,
		},
	}

	return commandMap
}
