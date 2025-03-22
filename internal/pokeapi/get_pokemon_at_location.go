package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemonAtLocation(locationId string) (pokemonsAtLocation, error) {
	var pokemons pokemonsAtLocation
	url := baseURL + "/location-area/" + locationId

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &pokemons); err != nil {
			return pokemons, nil
		}
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return pokemonsAtLocation{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return pokemonsAtLocation{}, err
	}

	if err := json.Unmarshal(data, &pokemons); err != nil {
		return pokemonsAtLocation{}, err
	}

	c.cache.Add(url, data)

	return pokemons, nil
}
