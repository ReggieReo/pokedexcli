package main

import (
	"time"

	"github.com/ReggieReo/pokedexcli/internal/pokeapi"
)
type config struct {
	client pokeapi.Client
	next *string
	previous *string
}

func main() {
	c := config{
		client: pokeapi.NewClient(time.Minute * 5),
	}
	run(&c)
}
