package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//HTTP fetch function 
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon" + "/" + name

	if data, found := c.cache.Get(url); found {
		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
	return pokemon, nil	
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

	c.cache.Add(url, body)
	return pokemon, nil
}



