package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex")
	availableCommands := getCommands()
	fmt.Println("Here are the available commands:")
	for _, cmd := range availableCommands {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
