package main

import (
	"fmt"
	"math/rand"
	"errors"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("Please add 1 pokemon to catch")
	}

	pokemonDetailsResp, err := cfg.pokeapiClient.GetPokemon(&args[0])
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonDetailsResp.Name)

	roll := rand.Intn(1000)
	threshold := 1000 - pokemonDetailsResp.BaseExperience*3

	if threshold < 0 {
		threshold = 50
	}

	if roll < threshold {
		cfg.PlayerPokemon[pokemonDetailsResp.Name] = pokemonDetailsResp
		fmt.Printf("%s was caught!\n", pokemonDetailsResp.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetailsResp.Name)
	}

	return nil
}