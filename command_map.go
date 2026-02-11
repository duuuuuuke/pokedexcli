package main

import (
	"errors"
	"fmt"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ApiResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func commandMapf(cfg *config, args ...string) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationRes.Next
	cfg.previousLocationsURL = &locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationRes.Next
	cfg.previousLocationsURL = &locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}
