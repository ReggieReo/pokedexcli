package main

import (
	"errors"
	"fmt"

	"github.com/ReggieReo/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	locationsINFO, err := c.client.GetLocation(c.next)
	if err != nil {
		return err
	}

	c.next = locationsINFO.Next
	c.previous = locationsINFO.Previous
	printLocation(locationsINFO)
	return nil
}

func commandMapb(c *config) error {
	if c.previous == nil {
		return errors.New("you are on the first page")
	}
	locationsINFO, err := c.client.GetLocation(c.previous)
	if err != nil {
		return err
	}
	c.next = locationsINFO.Next
	c.previous = locationsINFO.Previous
	printLocation(locationsINFO)
	return nil
}

func printLocation(mapAreas pokeapi.Location) {
	for _, area := range mapAreas.Results {
		fmt.Printf("-%v\n", area.Name)
	}

}
