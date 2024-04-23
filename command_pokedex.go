package main

import "fmt"

func callbackPokedex(cfg *config, _ ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("you have not catched any pokemons")
		return nil
	}
	fmt.Printf("Your Pokedex\n")
	for _, v := range cfg.pokedex {
		fmt.Printf("- %s\n", v.Name)
	}
	return nil
}
