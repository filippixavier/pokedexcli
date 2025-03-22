package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, params ...string) error {
	if len(params) == 0 {
		return errors.New("missing location argument")
	}

	pokemon, err := config.Client.GetPokemonInfo(params[0])

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])

	pokethrow := rand.Intn(pokemon.BaseExperience)

	if pokethrow < 40 {
		fmt.Printf("%s caugth!\n", pokemon.Name)
		config.Pokedex[pokemon.Name] = pokemon
		fmt.Println("You may now inspect if with the inspect command")
	} else {
		fmt.Println("Too bad...try again!")
	}
	return nil
}
