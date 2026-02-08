package main

import (
	"github.com/vetal-bla/bootdev-pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeApiClient: pokeClient,
	}
	startRepl(cfg)
}
