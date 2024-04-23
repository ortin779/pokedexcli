package main

import (
	"fmt"
)

func callbackExplore(cfg *config, options ...string) error {
	area := options[0]
	fmt.Printf("Exploring %s...", area)
	resp, err := cfg.pokeapiClient.GetLocationAreaInfo(&area)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}
	return nil
}
