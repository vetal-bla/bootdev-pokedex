package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func isCaught(exp int) bool {

	var chanceForCatch float64

	if exp > 0 && exp <= 50 {
		chanceForCatch = 0.9
	} else if exp > 51 && exp <= 100 {
		chanceForCatch = 0.7
	} else if exp > 101 {
		chanceForCatch = 0.5
	}

	return rand.Float64() < chanceForCatch
}

func catch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := c.pokeApiClient.PokemonInfo(pokemonName)
	if err != nil {
		fmt.Printf("Can't get pokemon info. Sorry...\n %s", err)
		return err
	}

	if isCaught(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonName)
		if _, ok := c.userDex[pokemonName]; !ok {
			c.userDex[pokemonName] = pokemon
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
