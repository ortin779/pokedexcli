package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to Pokedex help menu:")
	fmt.Println("Here are your available commands")

	commands := getAvailableCommands()

	for _, cmd := range commands {
		fmt.Printf("-%s : %s\n", cmd.name, cmd.description)
	}
	return nil
}
