package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

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
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect caught <pokemon> in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the pokemon in your pokedex",
			callback:    commandPokedex,
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
	pokemon, err := ConvertDetails(&pokeDetails)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	caught := isCaught(pokemon.BaseExperience)
	if caught {
		fmt.Printf("You caught %s! Storing in Pokedex...\n", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
		fmt.Println("Pokemon stored!")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	// fmt.Println(pokeDetails)

	return nil
}

func commandInspect(pokeName string) error {
	pokemon, ok := pokedex[pokeName]
	if !ok {
		fmt.Printf("You haven't caught %s yet...", pokeName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Printf("  -%s\n", pType)
	}
	return nil

}

func commandPokedex(_ string) error {
	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokexex:")
	for key := range pokedex {
		fmt.Printf("  -%s\n", key)
	}
	return nil
}

func isCaught(baseExp int) bool {
	var median int = 125
	var threshold = float32(baseExp) / float32(baseExp+median)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	val := r.Float32()

	// fmt.Println(baseExp, threshold, val)

	return val > float32(threshold)

}
