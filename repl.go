package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/filippixavier/pokedexcli/internal/pokeapi"
)

type Config struct {
	Client  pokeapi.Client
	Pokedex map[string]pokeapi.PokemonInfo
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	config := Config{
		Client:  pokeapi.NewClient(5 * time.Second),
		Pokedex: make(map[string]pokeapi.PokemonInfo),
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())

		if len(input) == 0 {
			continue
		}

		if cmd, exist := commands[input[0]]; exist {
			params := input[1:]
			if err := cmd.callback(&config, params...); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
