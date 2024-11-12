package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func CommandCatch(config *Config, pokemon string) error {

	fmt.Println()

	if _, ok := config.Pokemon[pokemon]; ok {
		return fmt.Errorf("%s is already in your pokedex", pokemon)
	}

	const baseURL = "https://pokeapi.co/api/v2/pokemon/"
	fullURL := baseURL + pokemon

	var data []byte

	data, inCache := config.Cache.Get(fullURL)

	if !inCache {
		// make new network request to get data
		res, err := http.Get(fullURL)
		if res.StatusCode == 404 {
			return fmt.Errorf("pokemon not found - ensure you are entering the pokemon name accurately")
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

	// Now we can proceed with unmarshaling the data
	var pokemonDetails PokemonDetailed
	if err := json.Unmarshal(data, &pokemonDetails); err != nil {
		return err
	}

	// Use Pokemon's BaseExperience stat plus random number generation to determine whether Pokemon is caught
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	baseExperience := pokemonDetails.BaseExperience
	cutoff := float64(baseExperience) * .75
	if rand.Intn(baseExperience) >= int(cutoff) {
		// successfully caught pokemon
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Printf("Adding %s to your pokedex...\n", pokemon)
		config.Pokemon[pokemon] = Pokemon{
			Name:   pokemonDetails.Name,
			Height: pokemonDetails.Height,
			Weight: pokemonDetails.Weight,
			Stats: []struct {
				BaseStat int
				Effort   int
				Stat     struct {
					Name string
					URL  string
				}
			}(pokemonDetails.Stats),
			Types: []struct {
				Slot int
				Type struct {
					Name string
					URL  string
				}
			}(pokemonDetails.Types),
		}

	} else {
		// Pokemon got away
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
