package main

import (
	"errors"
	"fmt"
)

func commandMapF(config *Config, _params ...string) error {
	locations, err := config.Client.GetLocations(config.Client.Next)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.Client.Next = locations.Next
	config.Client.Previous = locations.Previous
	return nil
}

func commandMapBack(config *Config, _params ...string) error {
	if config.Client.Previous == nil {
		return errors.New("you are on the first page")
	}
	locations, err := config.Client.GetLocations(config.Client.Previous)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.Client.Next = locations.Next
	config.Client.Previous = locations.Previous
	return nil
}
