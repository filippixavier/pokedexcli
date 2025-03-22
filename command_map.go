package main

import (
	"errors"
	"fmt"

	"github.com/filippixavier/pokedexcli/internal/pokeapi"
)

func commandMapF(config *pokeapi.Client, _params ...string) error {
	locations, err := config.GetLocations(config.Next)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}

func commandMapBack(config *pokeapi.Client, _params ...string) error {
	if config.Previous == nil {
		return errors.New("you are on the first page")
	}
	locations, err := config.GetLocations(config.Previous)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}
