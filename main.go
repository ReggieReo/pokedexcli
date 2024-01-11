package main

import "github.com/ReggieReo/pokedexcli/internal/pokeapi"

type config struct {
	client pokeapi.Client
	next *string
	previous *string
}

func main() {
	c := config{
		client: pokeapi.NewClient(),
	}
	run(&c)
}
