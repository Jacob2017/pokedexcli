package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonDetails struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
	} `json:"cries"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetPokemonDetails(pokeName string) (PokemonDetails, error) {
	var url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokeName)
	var pokemon PokemonDetails

	res, err := http.Get(url)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("Network error: %v", err)
	}

	defer res.Body.Close()

	// var pokemon PokemonDetails
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&pokemon); err != nil {
		return PokemonDetails{}, fmt.Errorf("error decoding response body: %v", err)
	}

	return pokemon, nil

}
