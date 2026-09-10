package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepel(cfg *config) {
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
		args := []string{}
		if len(cleanedWords) > 1 {
			args = cleanedWords[1:]
		}
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			fmt.Println("Enter 'help' to see the list of available commands.")
			continue
		}

		command.callback(cfg, args...)

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"map": {
			name:        "map",
			description: "Displays the name of the next location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of the previous location areas",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore {location_area_name}",
			description: "Displays the name of the pokemon in a location area",
			callback:    callbackExplore,
		},
	}
}

func cleanInput(input string) []string {
	lowerInput := strings.ToLower(input)
	words := strings.Fields(lowerInput)

	return words
}
