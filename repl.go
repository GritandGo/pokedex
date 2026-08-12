package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(cfg *config) {
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
		if cmd, ok := cfg.commandConfig[firstWord]; ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}


func cleanInput(text string) []string {
	splitInput := strings.Fields(text)
	for i := 0; i < len(splitInput); i++ {
		splitInput[i] = strings.ToLower(splitInput[i])
	}
	return splitInput
	
}
	