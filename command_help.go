package main

import "fmt"

func commandHelp(currentLocation *CurrentLocation) error {
    fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommandMap() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
    return nil
}
