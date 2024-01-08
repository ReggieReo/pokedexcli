package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, com := range getCommand() {
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	fmt.Println()
	return nil
}
