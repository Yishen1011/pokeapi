package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(currentLocation *CurrentLocation) error {
	url := currentLocation.Next
	maps, err := getMap(*url)
	if err != nil {
		return fmt.Errorf("Can't receive maps:  %w", err)
	}

	currentLocation.Next = maps.NextMap
	currentLocation.Previous = maps.PreviousMap
	
	err = listMaps(maps)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func commandMapb(currentLocation *CurrentLocation) error {
	url := currentLocation.Previous

	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	maps, err := getMap(*url)
	if err != nil {
		return fmt.Errorf("Can't receive maps:  %w", err)
	}

	currentLocation.Next = maps.NextMap
	currentLocation.Previous = maps.PreviousMap
	
	err = listMaps(maps)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func listMaps(maps Map) error {
	for _, location := range maps.Locations {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func getMap(url string) (Map, error) {
	res, err := http.Get(url)
	if err != nil {
		return Map{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Map{}, fmt.Errorf("error reading response: %w", err)
	}

	var locationMap Map
	if err := json.Unmarshal(data, &locationMap); err != nil {
		return Map{}, err
	}

	return locationMap, nil
}