package main

import (
	"os"

	"github.com/ajswetz/go-pokedex-cli/internal/pokeapi"
)

func commandExit(config *pokeapi.Config, _ string) error {
	os.Exit(0)
	return nil
}
