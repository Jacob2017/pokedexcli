package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jacob2017/pokedexcli/internal/pokecache"
)

var apiCache = pokecache.NewCache(5000 * time.Millisecond)

func GetLocationAreaBatch(offset int) ([]LocationArea, error) {
	var baseUrl = "https://pokeapi.co/api/v2/location-area?limit=20"
	var url = baseUrl + fmt.Sprintf("&offset=%d", offset)
	var areas []LocationArea

	if cachedVal, ok := apiCache.Get(url); ok {
		if err := json.Unmarshal(cachedVal, &areas); err != nil {
			return []LocationArea{}, fmt.Errorf("Error unmarshalling cached value: %v", err)
		}
		return areas, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return []LocationArea{}, fmt.Errorf("Network error: %v", err)
	}

	defer res.Body.Close()

	var result LocationAreaResult
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		return []LocationArea{}, fmt.Errorf("Error decoding response body: %v", err)
	}

	areas = result.Results

	cacheData, err := json.Marshal(areas)
	if err != nil {
		return nil, fmt.Errorf("error marshalling for cache: %v", err)
	}
	apiCache.Add(url, cacheData)

	return areas, nil
}

func GetLocationAreaDetails(locationName string) ([]PokeEncounters, error) {
	var url = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationName)
	var pokemon []PokeEncounters

	if cachedVal, ok := apiCache.Get(url); ok {
		if err := json.Unmarshal(cachedVal, &pokemon); err != nil {
			return []PokeEncounters{}, fmt.Errorf("Error unmarshalling cached value: %v", err)
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return []PokeEncounters{}, fmt.Errorf("Network error: %v", err)
	}

	defer res.Body.Close()

	var result LocationAreaDetails
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&result); err != nil {
		return []PokeEncounters{}, fmt.Errorf("Error decoding response body: %v", err)
	}

	pokemon = result.PokemonEncounters
	cacheData, err := json.Marshal(pokemon)
	if err != nil {
		return nil, fmt.Errorf("error marshalling for cache: %v", err)
	}
	apiCache.Add(url, cacheData)
	return pokemon, nil
}
