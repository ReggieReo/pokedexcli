package main

import "fmt"

func commandPokedex(c *config, input ...string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("You have not cauch any pokemon")
	return nil
	}

	fmt.Println("Your pokedex:")
	for _, poke:= range c.pokedex {
		fmt.Println(" - ", poke.Name)

	}
	return nil
}
