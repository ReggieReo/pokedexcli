package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaInfo(locationName string) (AreaPokemon, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint + locationName

	areaPokemons := AreaPokemon{}
	cData, ok := c.cache.Get(fullURL)
	// if there cache unmashal data and return them
	if ok {
		err := json.Unmarshal(cData, &areaPokemons)
		if err != nil {
			return AreaPokemon{}, err
		}
		return areaPokemons, nil
	}

	// else make a request normally
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return AreaPokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaPokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaPokemon{}, err
	}
	c.cache.Add(fullURL, data)
	err = json.Unmarshal(data, &areaPokemons)
	if err != nil {
		return AreaPokemon{}, err
	}
	return areaPokemons, nil
}
