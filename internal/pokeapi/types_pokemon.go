package pokeapi



type Pokemon struct {
	BaseExperience int `json:"base_experience"`
	Name string `json:"name"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
} `json:"stats"`
}



