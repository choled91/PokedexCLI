package main

import "github.com/choled91/PokedexCLI/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAReaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepel(&cfg)
}
