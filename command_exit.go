package main

import "os"

func callbackExit(cfg *config) error {
	os.Exit(0)
	return nil
}
