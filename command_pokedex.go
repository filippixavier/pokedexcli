package main

import "fmt"

func commandPokedex(config *Config, _params ...string) error {
	pokedex := &config.Pokedex

	if len(*pokedex) == 0 {
		fmt.Println("Your pokédex is empty, start catching pokémon using the catch command!")
		return nil
	}

	fmt.Println("Your Pokédex:")
	for k, _ := range *pokedex {
		fmt.Printf("	- %s\n", k)
	}

	return nil
}
