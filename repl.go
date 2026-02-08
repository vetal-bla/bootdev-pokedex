package main

import (
	"bufio"
	"fmt"
	"github.com/vetal-bla/bootdev-pokedex/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeApiClient pokeapi.Client
	nextUrl       *string
	previousUrl   *string
}

func startRepl(c *config) {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		userInputs := cleanInput(reader.Text())

		if len(userInputs) == 0 {
			continue
		}
		command := userInputs[0]

		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Println("unknown command")
		} else {
			if err := cmd.callback(c); err != nil {
				fmt.Println(err)
			}
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Disaplay next 20 locations",
			callback:    mapNames,
		},
		"mapb": {
			name:        "mapb",
			description: "Disaplay previous 20 locations",
			callback:    mapbNames,
		},
	}
}
