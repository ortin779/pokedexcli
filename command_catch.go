package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, options ...string) error {
	pokemonName := options[0]
	resp, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	randInt := rand.Intn(255)

	if resp.BaseExperience < randInt {
		fmt.Printf("%s catched!\n", pokemonName)
		cfg.pokedex[pokemonName] = resp
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}
