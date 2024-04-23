package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	val, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResp := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad request %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreaResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)
	return locationAreaResp, nil
}

func (c *Client) GetLocationAreaInfo(area *string) (LocationArea, error) {
	if area == nil || len(*area) == 0 {
		return LocationArea{}, fmt.Errorf("the entered area : %s is not a valid area", *area)
	}
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint + *area

	val, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad request %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResp := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)
	return locationAreaResp, nil
}

func (c *Client) GetPokemonInfo(pokemon string) (Pokemon, error) {
	if len(pokemon) == 0 {
		return Pokemon{}, fmt.Errorf("the entered area : %s is not a valid area", pokemon)
	}
	endpoint := "/pokemon/"
	fullURL := baseURL + endpoint + pokemon

	val, ok := c.cache.Get(fullURL)
	if ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad request %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)
	return pokemonResp, nil
}
