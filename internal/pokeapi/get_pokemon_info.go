package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemonInfo(pokemonName string) (PokemonInfo, error) {
	var pokemons PokemonInfo
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &pokemons); err != nil {
			return pokemons, nil
		}
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return PokemonInfo{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonInfo{}, err
	}

	if err := json.Unmarshal(data, &pokemons); err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, data)

	return pokemons, nil
}
