package main

import (
	"errors"
	"fmt"
)

func inspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name")
	}

	pokemonName := args[0]

	if v, ok := c.userDex[pokemonName]; ok {
		fmt.Printf(`
Name: %s
Height: %d
Weight: %d
Stats:
			`, v.Name, v.Height, v.Weight)
		for _, elem := range v.Stats {
			fmt.Printf("  - %s: %d\n", elem.Stat.Name, elem.BaseStat)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
