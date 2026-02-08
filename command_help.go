package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
