package main

import "github.com/ajswetz/go-pokedex-cli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 names of location areas in the Pokemon world",
			callback:    pokeapi.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 names of location areas in the Pokemon world",
			callback:    pokeapi.CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Shows a list of Pokemon available in a given location",
			callback:    pokeapi.CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon by name",
			callback:    pokeapi.CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Gets the details of a Pokemon that you've already caught",
			callback:    pokeapi.CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints a list of all Pokemon you have caught",
			callback:    pokeapi.CommandPokedex,
		},
	}

}
