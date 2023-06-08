package request

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	pokemonApI = "https://pokeapi.co/api/v2/ability/?offset=1&limit=1"
)

type pokemon struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func TestGet(t *testing.T) {
	c := require.New(t)

	client, err := NewHTTPClient()
	c.NoError(err)
	c.NotNil(client)

	body, err := client.Get(pokemonApI)
	c.NoError(err)

	var animal pokemon
	err = json.Unmarshal(body, &animal)
	c.NoError(err)

	c.NotEmpty(animal.Results[0].Name)
}

func TestGetEmpty(t *testing.T) {
	c := require.New(t)

	client, err := NewHTTPClient()
	c.NoError(err)
	c.NotNil(client)

	_, err = client.Get("")
	c.Error(err, ErrURLEmpty)
}
