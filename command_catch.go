package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, input ...string) error {
	pokemonName := input[1]
	pokemon, err := c.client.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	res := rand.Intn(pokemon.BaseExperience)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	c.pokedex[pokemon.Name] = pokemon
	return nil
}
