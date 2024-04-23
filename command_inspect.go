package main

import "fmt"

func callbackInspect(cfg *config, options ...string) error {
	pokemonName := options[0]

	if len(pokemonName) == 0 {
		return fmt.Errorf("in-valid pokemon name: %s", pokemonName)
	}

	val, ok := cfg.pokedex[pokemonName]
	if ok {
		fmt.Print(val)
		return nil
	}
	fmt.Println("you have not caught that pokemon")
	return nil

}
