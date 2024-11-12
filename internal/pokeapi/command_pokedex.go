package pokeapi

import (
	"fmt"
)

func CommandPokedex(config *Config, _ string) error {

	fmt.Println()

	if len(config.Pokemon) == 0 {
		return fmt.Errorf("pokedex empty - you haven't caught any Pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for key, _ := range config.Pokemon {
		fmt.Printf(" - %v\n", key)
	}

	return nil
}
