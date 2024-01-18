package main

import (
	"time"

	"github.com/ReggieReo/pokedexcli/internal/pokeapi"
)

type config struct {
	client   pokeapi.Client
	next     *string
	previous *string
	pokedex  map[string]pokeapi.Pokemon
}

func main() {
	c := config{
		client:  pokeapi.NewClient(time.Minute * 5),
		pokedex: map[string]pokeapi.Pokemon{},
	}
	run(&c)
}
