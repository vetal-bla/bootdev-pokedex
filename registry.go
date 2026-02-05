package main

import (
	"fmt"
	"github.com/vetal-bla/bootdev-pokedex/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	nextUrl     string
	previousUrl string
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

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func mapNames(c *config) error {
	mapAreas, err := pokeapi.GetLocationAreas(c.nextUrl)
	if err != nil {
		return err
	}
	if mapAreas.Next != "" {
		c.nextUrl = mapAreas.Next
	} else {
		fmt.Println("Last page nothing to fetch")
		return nil
	}
	c.previousUrl = mapAreas.Previous

	for _, area := range mapAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func mapbNames(c *config) error {
	if c.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	mapAreas, err := pokeapi.GetLocationAreas(c.previousUrl)

	if err != nil {
		fmt.Println("Error fetching resource")
		return err
	}
	c.nextUrl = mapAreas.Next
	c.previousUrl = mapAreas.Previous
	for _, area := range mapAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
