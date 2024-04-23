package main

import "os"

func callbackExit(cfg *config, options ...string) error {
	os.Exit(0)
	return nil
}
