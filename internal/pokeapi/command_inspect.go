package pokeapi

import (
	"fmt"
)

func CommandInspect(config *Config, pokemon string) error {

	if _, ok := config.Pokemon[pokemon]; !ok {
		return fmt.Errorf("cannot inspect %s - you must catch a Pokemon before you can inspect it", pokemon)
	}

	fmt.Printf("Name: %v\n", config.Pokemon[pokemon].Name)
	fmt.Printf("Height: %v\n", config.Pokemon[pokemon].Height)
	fmt.Printf("Weight: %v\n", config.Pokemon[pokemon].Weight)

	fmt.Println("Stats:")
	fmt.Printf(" -hp: %v\n", config.Pokemon[pokemon].Stats[0].BaseStat)
	fmt.Printf(" -attack: %v\n", config.Pokemon[pokemon].Stats[1].BaseStat)
	fmt.Printf(" -defense: %v\n", config.Pokemon[pokemon].Stats[2].BaseStat)
	fmt.Printf(" -special-attack: %v\n", config.Pokemon[pokemon].Stats[3].BaseStat)
	fmt.Printf(" -special-defense: %v\n", config.Pokemon[pokemon].Stats[4].BaseStat)
	fmt.Printf(" -speed: %v\n", config.Pokemon[pokemon].Stats[5].BaseStat)

	fmt.Println("Types:")
	//Need a for loop here - each Pokemon might have a different amount of Types
	for _, slot := range config.Pokemon[pokemon].Types {
		fmt.Printf(" - %v\n", slot.Type.Name)
	}

	return nil
}
