package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, params ...string) error {
	if len(params) == 0 {
		return errors.New("missing pokemon name")
	}

	pokedex := &config.Pokedex

	if pkmn, ok := (*pokedex)[params[0]]; !ok {
		return fmt.Errorf("%s has not been caught yet", params[0])
	} else {
		fmt.Printf("Name: %v\n", pkmn.Name)
		fmt.Printf("Height: %v\n", pkmn.Height)
		fmt.Printf("Weight: %v\n", pkmn.Weight)
		fmt.Println("Stats:")
		for _, stat := range pkmn.Stats {
			fmt.Printf("	-%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pType := range pkmn.Types {
			fmt.Printf("	-%v\n", pType.Type.Name)
		}
	}
	return nil
}
