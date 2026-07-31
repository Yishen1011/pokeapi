package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
    fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
    return nil
}
