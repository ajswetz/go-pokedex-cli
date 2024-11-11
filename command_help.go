package main

import (
	"fmt"

	"github.com/ajswetz/go-pokedex-cli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Config, _ string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
