package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepel() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleanedWords := cleanInput(text)
		if len(cleanedWords) == 0 {
			continue
		}
		commandName := cleanedWords[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			fmt.Println("Enter 'help' to see the list of available commands.")
			continue
		}

		command.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
	}
}

func cleanInput(input string) []string {
	lowerInput := strings.ToLower(input)
	words := strings.Fields(lowerInput)

	return words
}
