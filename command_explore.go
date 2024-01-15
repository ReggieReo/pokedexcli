package main

import "fmt"

func commandExplore(c *config, input ...string) error {
	if len(input) != 2 {
		return fmt.Errorf("the len of input is not two, %v", len(input))
	}
	areaName := input[1]
	pokes, err := c.client.GetLocationAreaInfo(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon: ")
	for _, enc := range pokes.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
