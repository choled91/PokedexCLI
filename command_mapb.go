package main

import (
	"fmt"
)

func callbackMapB(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAReaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
