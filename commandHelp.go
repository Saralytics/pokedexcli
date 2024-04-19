package main

import "fmt"

func callbackHelp(config *config) error {
	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Println(cmd.name)
		fmt.Println(cmd.description)
		fmt.Println("")
	}
	return nil
}
