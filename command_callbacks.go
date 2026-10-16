package main
import (
	"fmt"
	"os"
	"net/http"
	"io"
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
	res, err := http.Get("http://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	
	defer res.Body.Close() 

	body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		fmt.Println(string(body))
		}
	return nil
}