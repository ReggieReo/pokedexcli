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
	callback    func(*config, ...string) error
}


func run(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
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
		err := cCommand.callback(c, cleaned...)
		if err != nil {
			fmt.Println(err)
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "explore the given the map area",
			callback: commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "inspect catched pokemon",
			callback: commandInspect,
		},
		
	}
	return command
}
