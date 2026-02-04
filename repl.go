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

		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Println("unknown command")
		} else {
			if err := cmd.callback(); err != nil {
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
