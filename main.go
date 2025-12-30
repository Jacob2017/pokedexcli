package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input = scanner.Text()
			cleaned := cleanInput(input)
			// fmt.Print("\n")
			// fmt.Printf("Your command was: %s\n", cleaned[0])
			regFunc, ok := registry[cleaned[0]]
			if ok {
				var args string
				if len(cleaned) > 1 {
					args = strings.Join(cleaned[1:], " ")
				} else {
					args = ""
				}
				regFunc.callback(args)

			} else {
				fmt.Println("Unknown command")
			}

			// keep this at the bottom
			fmt.Print("Pokedex > ")
		}
	}
}
