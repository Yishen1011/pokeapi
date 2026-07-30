package main

import "fmt"


func commandExplore(cfg *config) error {
	encountersResp, err := cfg.pokeapiClient.ListPokemons(cfg.locationArea)
	if err != nil {
		return err
	}
	
	fmt.Printf("Exploring %s...\n", *cfg.locationArea)
	for _, encounters := range encountersResp.Encounters {
		fmt.Printf("%s\n", encounters.Pokemon.Name)
	}

	return nil
}