package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.PlayerPokemon) == 0 {
		fmt.Printf("You haven't caught any pokemons\n")
	} else {
		fmt.Printf("Your Pokedex:\n")
		for _, pokemon := range cfg.PlayerPokemon {
			fmt.Printf("\t- %s\n", pokemon.Name)
		}
	}

	return nil
}