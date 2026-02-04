package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		userInputs := cleanInput(reader.Text())

		if len(userInputs) == 0 {
			continue
		}
		command := userInputs[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
