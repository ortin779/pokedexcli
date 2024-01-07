package main

import "github.com/ortin779/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextPageURL   *string
	prevPageURL   *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		nextPageURL:   nil,
		prevPageURL:   nil,
	}
	repl(&cfg)
}
