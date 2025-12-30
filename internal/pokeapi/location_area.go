package pokeapi

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResult struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationAreaDetails struct {
	Name              string           `json:"name"`
	PokemonEncounters []PokeEncounters `json:"pokemon_encounters"`
}

type PokeEncounters struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}
