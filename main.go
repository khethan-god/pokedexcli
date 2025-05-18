package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var output []string

	// trim any leading or trailing space
	input := strings.TrimSpace(text)

	// convert the string to lower case
	lowerStr := strings.ToLower(input)

	// split the string
	output = strings.Split(lowerStr, " ")

	// Filter out empty strings resulting from multiple spaces
	var cleanedOutput []string
	for _, word := range output {
		if word != "" {
			cleanedOutput = append(cleanedOutput, word)
		}
	}
	return cleanedOutput
}
