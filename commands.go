package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands *map[string]cliCommand) error {
	fmt.Println("Usage:")
	fmt.Println()

	for _, val := range *commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	var commands map[string]cliCommand
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit de the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback: func() error {
				return commandHelp(&commands)
			},
		},
	}
	return commands
}
