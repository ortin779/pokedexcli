package main

import "fmt"

func callbackMapb(cfg *config) error {
	if cfg.prevPageURL == nil {
		return fmt.Errorf("error: can't go back, you reached first page\n")
	}

	locations, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevPageURL)

	if err != nil {
		return err
	}

	fmt.Println("List of locations :")
	for _, location := range locations.Results {
		fmt.Printf(" -%s", location.Name)
	}
	return nil
}
