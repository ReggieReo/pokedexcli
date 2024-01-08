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
	callback    func() error
}

func run() {
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
		err := cCommand.callback()
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
	}
	return command
}