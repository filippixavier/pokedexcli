package main

import (
	"fmt"
	"os"

	"github.com/filippixavier/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.Client) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
