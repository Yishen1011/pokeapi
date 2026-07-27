package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)


	url := "https://pokeapi.co/api/v2/location-area"
	currentLocation := &CurrentLocation{
		Next: &url,
		Previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command, exists := getCommandMap()[words[0]]
        if exists {
			err := command.callback(currentLocation)
			if err != nil {
				fmt.Println(err)
			}
        } else {
			fmt.Println("Unknown command")
            continue
        }
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*CurrentLocation) error
}

func getCommandMap() map[string]cliCommand{
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
			description: "Displays all map in current location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays all previous map in current location",
			callback:    commandMapb,
		},
	}
}