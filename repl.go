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
	callback    func(cfg *config) error
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

		err := command.callback(cfg)
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
