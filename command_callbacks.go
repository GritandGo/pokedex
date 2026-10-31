package main
import (
	"fmt"
	"os"
	"errors"
	"math/rand"
)


func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	for _, cmd := range cfg.commandConfig {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	 }

	return nil
	}


func commandMap(cfg *config, args []string) error {
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

func commandMapb(cfg *config, args []string) error {
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


func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("enter a valid area to explore")
	}

	cityName := args[0]
	location, err := cfg.pokeapiClient.GetLocation(cityName)
	if err != nil {
		return err
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}


func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("enter a valid Pokémon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	multiplier := 50.0
	successChance := multiplier / float64(pokemon.BaseExperience)


	if successChance > 1.0 {
		successChance = 1.0
	}

	if rand.Float64() < successChance {
		cfg.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

