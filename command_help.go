package main

import (
	"fmt"

	"github.com/filippixavier/pokedexcli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Client) error {
	fmt.Println("Usage:")
	fmt.Println()

	for _, val := range getCommands() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}
