package main

import (
	"time"

	"github.com/GritandGo/pokedex/internal/pokeapi"
)


func main() {
	cfg := &config{
		commandConfig: getCommands(),
		pokeapiClient: pokeapi.NewClient(5 * time.Second, 5 * time.Minute),
	}

	startRepl(cfg)
}