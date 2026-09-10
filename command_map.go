package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAReaURL)
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
