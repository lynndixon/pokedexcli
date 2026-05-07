package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type locationResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type apiResponse struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []locationResult `json:"results"`
}

func commandMap(cfg *config) error {
	url := baseURL + "/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var resp apiResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *cfg.Previous

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var resp apiResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil

}
