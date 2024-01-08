package main

import (
	"fmt"

	"github.com/ReggieReo/pokedexcli/pokeapi"
)

func commandMap(c *config) error {
	if c.next == "" {
		return fmt.Errorf("next url is nil")
	}
	next, previous := pokeapi.GetLocation(c.next)
	c.next = next
	c.previous = previous
	return nil
}

func commandMapb(c *config) error {
	if c.previous == "" {
		return fmt.Errorf("previous url is nil")
	}
	next, previous := pokeapi.GetLocation(c.previous)
	c.next = next
	c.previous = previous
	return nil
}
