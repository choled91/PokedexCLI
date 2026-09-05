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
		fmt.Println("Enter text to echo: ")

		scanner.Scan()
		text := scanner.Text()

		fmt.Println("echoing: ", text)
	}
}

func cleanInput(input string) []string {
	lowerInput := strings.ToLower(input)
	words := strings.Fields(lowerInput)

	return words
}
