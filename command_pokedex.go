package main

import (
	"fmt"
)

func pokedex(c *config, args ...string) error {

	dex := c.userDex

	if len(dex) == 0 {
		fmt.Println("you dont have any pokemon in pokdex")
	} else {
		fmt.Println("Your pokedex:")
		for k, _ := range dex {
			fmt.Printf("  - %s\n", k)
		}
	}

	return nil
}
