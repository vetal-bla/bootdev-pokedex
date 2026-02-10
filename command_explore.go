package main

import (
	"errors"
	"fmt"
)

func explore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	name := args[0]

	exploredLocation, err := c.pokeApiClient.ExploreLocations(name)
	// fmt.Printf("explore location: %s\n", args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploting %s...\n", name)
	fmt.Println("Found Pokemons:")

	for _, pokemon := range exploredLocation.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)

	}

	return nil
}
