package main

import "github.com/filippixavier/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Client, ...string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit de the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "display the next 20 locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "display the previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "display the list of pokemon that can be encouter at the given location",
			callback:    commandExploreLocation,
		},
	}
	return commands
}
