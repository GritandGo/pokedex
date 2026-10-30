package main
import (
	"github.com/GritandGo/pokedex/internal/pokeapi"
)

type config struct {
	commandConfig map[string]cliCommand
	Next *string
	Previous *string
	pokeapiClient pokeapi.Client
}


type cliCommand struct {
	name string
	description string
	callback func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},

	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},

	"map": {
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback:    commandMap,
	},

	"mapb": {
		name:        "mapb",
		description: "Displays 20 previous location areas in the Pokemon world",
		callback:    commandMapb,
	},

	"explore": {
		name:        "explore",
		description: "Displays the names of all Pokémon found in a specific area",
		callback:    commandExplore, 
	},
  }
}
