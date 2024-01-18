package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	cData, ok := c.cache.Get(fullURL)
	// if there cache unmashal data and return them
	if ok {
		err := json.Unmarshal(cData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullURL, dat)
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
