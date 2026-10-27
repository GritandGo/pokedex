package main
import (
	"fmt"
	"os"
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
	locationList, err := cfg.pokeapiClient.ListLocations(cfg.Next)
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
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} 


	locationList, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
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