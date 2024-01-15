package main

import (
	"os"
)

func commandExit(c *config, input ...string) error {
	os.Exit(0)
	return nil
}
