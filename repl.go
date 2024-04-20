package main

import (
	"bufio"
	"fmt"
	"os"
)

type clicommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

func startRepl(config *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Podedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		commandName := text[0]
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}
