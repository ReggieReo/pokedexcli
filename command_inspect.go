package main

import (
	"fmt"

	"github.com/ReggieReo/pokedexcli/internal/pokeapi"
)

func commandInspect(c *config, input ...string) error {
	pokeName := input[1]
	pokemon, ok := c.pokedex[pokeName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	printPokeinfo(pokemon)
	return nil
}

func printPokeinfo(pokemon pokeapi.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
}
