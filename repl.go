package main
import "strings"

func cleanInput(text string) []string {
	splitInput := strings.Fields(text)
	for i := 0; i < len(splitInput); i++ {
		splitInput[i] = strings.ToLower(splitInput[i])
	}
	return splitInput
	
}
	