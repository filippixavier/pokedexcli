package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())
		fmt.Printf("Your command was: %v\n", input[0])
	}
}
