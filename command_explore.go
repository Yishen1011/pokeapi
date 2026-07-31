package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("Please add the location to explore")
	}

	encountersResp, err := cfg.pokeapiClient.ListPokemons(&args[0])
	if err != nil {
		return err
	}
	
	fmt.Printf("Exploring %s...\n", args[0])
	for _, encounters := range encountersResp.Encounters {
		fmt.Printf("%s\n", encounters.Pokemon.Name)
	}

	return nil
}