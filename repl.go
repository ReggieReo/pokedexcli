package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next string
	previous string
}

func run() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		next: "https://pokeapi.co/api/v2/location/?limit=20&offset=20",
		previous: "",
	}
	for {
		// scan input
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		// run command
		cCommand, ok := getCommand()[cleaned[0]]
		// if command not exist continue to next iter
		if !ok {
			fmt.Println("The wrong command")
			continue
		}
		// else execute commandtext
		err := cCommand.callback(&c)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func cleanInput(str string) []string {
	str = strings.ToLower(str)
	return strings.Fields(str)
}

func getCommand() map[string]cliCommand {
	command := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 locations.",
			callback: commandMapb,
		},
	}
	return command
}
