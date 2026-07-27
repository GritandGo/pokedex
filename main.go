package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		cleanedInput := strings.Fields(input)
		
		if len(cleanedInput) == 0 {
			continue
		}

		firstWord := cleanedInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
