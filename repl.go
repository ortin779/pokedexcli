package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getAvailableCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints help manual for Pokedex",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex repl",
			callback:    callbackExit,
		},
	}
}

func repl() {
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

		command.callback()
	}
}

func cleanInput(str string) []string {
	lowerStr := strings.ToLower(str)
	words := strings.Fields(lowerStr)
	return words
}
