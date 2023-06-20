package request

import (
	"encoding/json"
	"testing"

	"github.com/MauCastillo/alana/shared/cnn/models"
	"github.com/stretchr/testify/require"
)

const (
	pokemonAPI = "https://pokeapi.co/api/v2/ability/?offset=1&limit=1"
	cnnAPI     = "https://production.dataviz.cnn.io/index/fearandgreed/graphdata/2023-06-12"
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

	body, err := client.Get(pokemonAPI)
	c.NoError(err)

	var animal pokemon
	err = json.Unmarshal(body, &animal)
	c.NoError(err)

	c.NotEmpty(animal.Results[0].Name)
}

func TestGetHeaders(t *testing.T) {
	c := require.New(t)

	client, err := NewHTTPClient()
	c.NoError(err)
	c.NotNil(client)

	header := []Header{{Key: "User-Agent", Value: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0"},
		{Key: "Accept", Value: "*/*"},
		{Key: "Accept-Language", Value: "en-US,en;q=0.5"},
		{Key: "Sec-Fetch-Dest", Value: "empty"},
		{Key: "Sec-Fetch-Mode", Value: "cors"},
		{Key: "Sec-Fetch-Site", Value: "cross-site"},
	}

	body, err := client.GetwithHeaders(cnnAPI, header)

	c.NoError(err)

	var response models.APIResponse
	err = json.Unmarshal(body, &response)
	c.NoError(err)

	c.Equal("extreme greed", response.FearAndGreed.Rating)
}
func TestGetEmpty(t *testing.T) {
	c := require.New(t)

	client, err := NewHTTPClient()
	c.NoError(err)
	c.NotNil(client)

	_, err = client.Get("")
	c.Error(err, ErrURLEmpty)
}
