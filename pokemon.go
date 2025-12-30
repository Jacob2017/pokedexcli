package main

import "github.com/jacob2017/pokedexcli/internal/pokeapi"

type Pokemon struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	BaseExperience int      `json:"base_experience"`
	CryURL         string   `json:"cry_url"`
	Types          []string `json:"types"`
	Stats          []Stat   `json:"stats"`
	Height         int      `json:"height"`
	Weight         int      `json:"weight"`
}

type Stat struct {
	Name     string `json:"name"`
	BaseStat int    `json:"base_stat"`
}

func ConvertDetails(pd *pokeapi.PokemonDetails) (Pokemon, error) {
	var newTypes = []string{}
	for _, pType := range pd.Types {
		newTypes = append(newTypes, pType.Type.Name)
	}
	var newStats = []Stat{}
	var newStat = Stat{}
	for _, stat := range pd.Stats {
		newStat = Stat{
			Name:     stat.Stat.Name,
			BaseStat: stat.BaseStat,
		}
		newStats = append(newStats, newStat)
	}
	var newPoke = Pokemon{
		ID:             pd.ID,
		Name:           pd.Name,
		BaseExperience: pd.BaseExperience,
		CryURL:         pd.Cries.Latest,
		Types:          newTypes,
		Stats:          newStats,
		Height:         pd.Height,
		Weight:         pd.Weight,
	}

	return newPoke, nil
}
