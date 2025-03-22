package main

import (
	"errors"
	"fmt"

	"github.com/filippixavier/pokedexcli/internal/pokeapi"
)

func commandExploreLocation(config *pokeapi.Client, params ...string) error {
	if len(params) == 0 {
		return errors.New("missing location argument")
	}

	pokemons, err := config.GetPokemonAtLocation(params[0])

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokémons:\n", params[0])

	for _, encouter := range pokemons.PokemonEncounters {
		fmt.Println(encouter.Pokemon.Name)
	}

	return nil
}
