package main

import "github.com/ortin779/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextPageURL   *string
	prevPageURL   *string
	pokedex       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		nextPageURL:   nil,
		prevPageURL:   nil,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	repl(&cfg)
}
