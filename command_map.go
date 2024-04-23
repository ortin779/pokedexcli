package main

import (
	"fmt"
)

func callbackMap(cfg *config, options ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextPageURL)
	if err != nil {
		return err
	}
	fmt.Println("List of locations:")
	for _, location := range resp.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	cfg.nextPageURL = resp.Next
	cfg.prevPageURL = resp.Previous
	return nil
}
