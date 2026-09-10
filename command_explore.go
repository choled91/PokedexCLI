package main

import (
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no location area name provided")
	}
	locattionAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locattionAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
