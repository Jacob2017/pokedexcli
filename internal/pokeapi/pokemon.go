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
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

func GetPokemonDetails(pokeName string) (PokemonDetails, error) {
	var url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokeName)
	var pokemon PokemonDetails

	if cachedVal, ok := apiCache.Get(url); ok {
		if err := json.Unmarshal(cachedVal, &pokemon); err != nil {
			return PokemonDetails{}, fmt.Errorf("Error unmarshalling cached value: %v", err)
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("Network error: %v", err)
	}

	defer res.Body.Close()

	// var pokemon Pokemon
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&pokemon); err != nil {
		return PokemonDetails{}, fmt.Errorf("error decoding response body: %v", err)
	}

	cacheData, err := json.Marshal(pokemon)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error marshalling for cache: %v", err)
	}
	apiCache.Add(url, cacheData)

	return pokemon, nil
}
