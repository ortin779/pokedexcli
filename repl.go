package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, options ...string) error
}

func getAvailableCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints help manual for Pokedex",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List some locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous locations",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the area for the available pokemons",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "cathc ${pokemon_name}",
			description: "Trying to catch a particular pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspecting the catched pokemons",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "printin the catched pokemons",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex repl",
			callback:    callbackExit,
		},
	}
}

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		cliCommands := getAvailableCommands()

		command, ok := cliCommands[cleaned[0]]

		if !ok {
			fmt.Println("Invalid Command")
			continue
		}

		err := command.callback(cfg, cleaned[1:]...)
		if err != nil {
			log.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowerStr := strings.ToLower(str)
	words := strings.Fields(lowerStr)
	return words
}
