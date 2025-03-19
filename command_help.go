package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Usage:")
	fmt.Println()

	for _, val := range getCommands() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}
