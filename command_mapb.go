package main

import "fmt"

func callbackMapb(cfg *config) error {
	if cfg.prevPageURL == nil {
		return fmt.Errorf("error: can't go back, you reached first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevPageURL)

	if err != nil {
		return err
	}

	fmt.Println("List of locations :")
	for _, location := range resp.Results {
		fmt.Printf(" -%s\n", location.Name)
	}
	cfg.nextPageURL = resp.Next
	cfg.prevPageURL = resp.Previous
	return nil
}
