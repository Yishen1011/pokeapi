package main

import (
	"fmt"
	"errors"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("Please add 1 pokemon to inspect")
	}

	pokemonDetail, ok := cfg.PlayerPokemon[args[0]]
	if ok {
		fmt.Printf("Name: %s\n", pokemonDetail.Name)
		fmt.Printf("Height: %d\n", pokemonDetail.Height)
		fmt.Printf("Weight: %d\n", pokemonDetail.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemonDetail.Stats {
			fmt.Printf("\t-%s: %d\n", stat.NameStat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, stat := range pokemonDetail.Types {
			fmt.Printf("\t- %s\n", stat.NameType.Name)
		}
	} else {
		fmt.Printf("you have not caught %s\n", pokemonDetail.Name)
	}

	return nil
}