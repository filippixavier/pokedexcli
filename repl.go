package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())

		if len(input) == 0 {
			continue
		}

		if cmd, exist := commands[input[0]]; exist {
			cmd.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
