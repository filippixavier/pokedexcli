package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/filippixavier/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	config := pokeapi.NewClient(5 * time.Second)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())

		if len(input) == 0 {
			continue
		}

		if cmd, exist := commands[input[0]]; exist {
			if err := cmd.callback(&config); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
