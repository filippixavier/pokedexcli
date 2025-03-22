package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocations(pageURL string) (locationsResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != "" {
		url = pageURL
	}

	fmt.Printf("Url: %s", url)

	res, err := c.httpClient.Get(url)

	if err != nil {
		return locationsResponse{}, fmt.Errorf("error when fetching next locations: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var locations locationsResponse

	if err := decoder.Decode(&locations); err != nil {
		return locationsResponse{}, fmt.Errorf("error when parsing response: %w", err)
	}

	return locations, nil
}
