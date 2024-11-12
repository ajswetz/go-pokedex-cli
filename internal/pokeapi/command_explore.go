package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CommandExplore(config *Config, location string) error {

	fmt.Println()

	const baseURL = "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL + location

	var data []byte

	data, inCache := config.Cache.Get(fullURL)

	if !inCache {
		// make new network request to get data
		res, err := http.Get(fullURL)
		if res.StatusCode == 404 {
			return fmt.Errorf("location not found - ensure you are entering the location name accurately")
		} else if err != nil {
			return err
		}

		// read http response into `data` variable as []byte
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// add this []byte data to the cache
		config.Cache.Add(config.Next, data)
	}

	// Now we can proceed with unmarshaling the data and printing it to the user
	var locAreaDetails LocationAreaDetails
	if err := json.Unmarshal(data, &locAreaDetails); err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locAreaDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
