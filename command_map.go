package main

import (
	"errors"
	"fmt"
)

func mapNames(c *config, args ...string) error {
	mapAreas, err := c.pokeApiClient.GetLocationAreas(c.nextUrl)
	if err != nil {
		return err
	}

	c.nextUrl = mapAreas.Next
	c.previousUrl = mapAreas.Previous

	for _, area := range mapAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func mapbNames(c *config, args ...string) error {
	if c.previousUrl == nil {
		return errors.New("you're on first page")
	}
	mapAreas, err := c.pokeApiClient.GetLocationAreas(c.previousUrl)

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
