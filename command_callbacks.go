package main
import (
	"fmt"
	"os"
	"net/http"
	"io"
	"github.com/GritandGo/pokedex/internal/pokeapi"
	"encoding/json"
)


func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	for _, cmd := range cfg.commandConfig {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	 }

	return nil
	}


func commandMap(cfg *config) error {
	var url string
	if cfg.Next == nil {
		url = "http://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	
	defer res.Body.Close() 

	body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var locationList pokeapi.LocationList
		err = json.Unmarshal(body, &locationList)
		if err != nil {
			return err
		}

		for _, area := range locationList.Results {
			fmt.Println(area.Name)
		}


		cfg.Next = locationList.Next
		cfg.Previous = locationList.Previous



	return nil
}


func commandMapb(cfg *config) error {
	var url string
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url = *cfg.Previous
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	
	defer res.Body.Close() 

	body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var locationList pokeapi.LocationList
		err = json.Unmarshal(body, &locationList)
		if err != nil {
			return err
		}

		for _, area := range locationList.Results {
			fmt.Println(area.Name)
		}


		cfg.Next = locationList.Next
		cfg.Previous = locationList.Previous



	return nil
}