package main

import (
	"bufio"
	"fmt"
	"os"
)

type clicommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

func startRepl(config *Config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Podedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		commandName := text[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		// if config.Next == nil {
		// 	fmt.Println("from repl:", config.Next)
		// } else {
		// 	fmt.Println("from repl:", *config.Next)
		// }

		if err := command.callback(config); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}

}
