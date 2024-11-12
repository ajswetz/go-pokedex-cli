package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CommandMapB(config *Config, _ string) error {

	fmt.Println()

	if config.Previous == nil {
		return fmt.Errorf("no previous location available - try using the `map` command first")
	}

	var data []byte

	data, inCache := config.Cache.Get(*config.Previous)

	if !inCache {
		// make new network request to get data
		res, err := http.Get(*config.Previous)
		if err != nil {
			return err
		}

		// read http response into `data` variable as []byte
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// add this []byte data to the cache
		config.Cache.Add(*config.Previous, data)

	}

	// Now we can proceed with unmarshaling the data and printing it to the user
	var locArea LocationAreaList
	if err := json.Unmarshal(data, &locArea); err != nil {
		return err
	}

	config.Next = locArea.Next
	config.Previous = locArea.Previous

	for _, location := range locArea.Results {
		fmt.Println(location.Name)
	}

	return nil
}
