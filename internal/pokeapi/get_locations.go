package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocations(pageURL *string) (locationsResponse, error) {
	var locations locationsResponse

	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &locations); err != nil {
			return locationsResponse{}, err
		}
		return locations, nil
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return locationsResponse{}, fmt.Errorf("error when fetching next locations: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return locationsResponse{}, err
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return locationsResponse{}, err
	}

	c.cache.Add(url, data)

	return locations, nil
}
