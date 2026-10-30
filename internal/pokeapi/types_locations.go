package pokeapi


type LocationList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []LocationArea `json:"results"`
	}


type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


type Location struct {
	PokemonEncounters []struct{
		Pokemon struct{
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}




