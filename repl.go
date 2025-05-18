package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	// convert the string to lower case
	output := strings.ToLower(text)
	// trim any leading or trailing space, splits them based on spaces and returns valid strings i.e., no spaces
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	// a registry of commands which can be used in our cli
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	/*
		this blocks the code and waits for user input from stdin, once the user
		types something and presses enter, .Scan() function returns true, and .Text()
		function stores the input, and the code continues
	*/
	for { // infinite loop
		fmt.Print("Pokedex > ")
		reader.Scan() // scans a line from stdin (console)

		words := cleanInput(reader.Text()) // pass the scaned text to cleanInput() to clean the input
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName] // selecting the command from a map
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
