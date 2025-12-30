package main

import (
	"fmt"
	"os"

	"github.com/jacob2017/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}

var registry map[string]cliCommand

func init() {
	registry = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Maps the world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Maps back one page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Explore <area_name> in more detail",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch <pokemon> with a poke ball",
			callback:    commandCatch,
		},
	}

}

func commandExit(_ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, cmd := range registry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(_ string) error {
	locations, err := pokeapi.GetLocationAreaBatch(currentMapOffset)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(locations) == 0 {
		fmt.Println("You're on the last page. Use `bmap` to go back.")
		return nil
	}

	currentMapOffset += 20
	for _, locArea := range locations {
		fmt.Println(locArea.Name)
	}

	return nil
}

func commandMapb(_ string) error {
	if currentMapOffset == 0 {
		fmt.Println("You're on the first page")
		return nil
	}

	currentMapOffset -= 20
	locations, err := pokeapi.GetLocationAreaBatch(currentMapOffset)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, locArea := range locations {
		fmt.Println(locArea.Name)
	}
	return nil
}

func commandExplore(areaName string) error {
	pokemon, err := pokeapi.GetLocationAreaDetails(areaName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(pokemon) > 0 {
		fmt.Println("Found Pokemon:")
	} else {
		fmt.Println("No Pokemon found!!")
	}

	for _, poke := range pokemon {
		fmt.Printf("- %s\n", poke.Pokemon.Name)
	}
	return nil
}

func commandCatch(pokeName string) error {
	pokeDetails, err := pokeapi.GetPokemonDetails(pokeName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeDetails.Name)
	fmt.Println(pokeDetails)

	return nil
}
