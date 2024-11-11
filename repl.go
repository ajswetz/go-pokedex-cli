package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	// initiate scanner on standard input
	replScanner := bufio.NewScanner(os.Stdin)

	// Use an infinite for loop to keep the REPL running.
	for {
		// At the start of the loop, you should block and wait for some input.
		fmt.Print("Pokedex > ")
		replScanner.Scan()

		// Once input is received, you should parse it and then execute a command.
		words := cleanInput(replScanner.Text())
		if len(words) == 0 {
			// blank input - continue to next loop iteration
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if !exists {
			// Unknown command - notify user and continue to next loop iteration
			fmt.Println("Unknown command")
			continue
		} else {
			// Legitimate command - process and then continue to next loop iteration
			err := command.callback()
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
