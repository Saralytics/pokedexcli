package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	words := strings.Fields(lower)

	return words
}

type clicommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp(commandMap map[string]clicommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range commandMap {
		fmt.Println(value.name, ": ", value.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Command not found. Run 'help' to see available commands.")
	os.Exit(0)
	return nil
}

func startRepl() {
	commandMap := map[string]clicommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
		},
		"map": {
			name:        "map",
			description: "display the next 20 locations",
		},
		"mapb": {
			name:        "mapb",
			description: "display the previous 20 locations",
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Podedex > ")
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		if _, ok := commandMap[text]; !ok {
			commandExit()
		}

		if text == "help" {
			commandHelp(commandMap)
			fmt.Print("Podedex > ")
		}

		if text == "exit" {
			os.Exit(0)
		}

	}

}
