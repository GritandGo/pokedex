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