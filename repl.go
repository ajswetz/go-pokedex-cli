package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ajswetz/go-pokedex-cli/internal/pokeapi"
	"github.com/ajswetz/go-pokedex-cli/internal/pokecache"
	"github.com/peterh/liner"
)

func startRepl() {

	// Build starting Config struct to pass into commands
	cmdConfig := pokeapi.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: nil,
		Cache:    *pokecache.NewCache(time.Second * 60),
		Pokemon:  make(map[string]pokeapi.Pokemon),
	}

	// Liner stuff
	line := liner.NewLiner()
	defer line.Close()

	// Use an infinite for loop to keep the REPL running.
	for {
		// At the start of the loop, you should block and wait for some input.
		input, err := line.Prompt("Pokedex > ")
		if err != nil {
			fmt.Printf("Error reading line: %v\n", err)
		}
		line.AppendHistory(input)

		// Once input is received, you should parse it and then execute a command.
		words := cleanInput(input)
		if len(words) == 0 {
			// blank input - continue to next loop iteration
			continue
		}

		commandName := words[0]

		// Initialize stringArg variable as empty string
		stringArg := ""
		// If an argument was passed in with the base command, set stringArg to that provided text
		if len(words) > 1 {
			stringArg = words[1]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			// Unknown command - notify user and continue to next loop iteration
			fmt.Println("Unknown command")
			continue
		} else {
			// Legitimate command - process and then continue to next loop iteration
			err := command.callback(&cmdConfig, stringArg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
