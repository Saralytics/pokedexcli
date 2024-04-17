package main

import (
	"os"
)

func callbackExit(config *Config) error {
	os.Exit(0)
	return nil
}
