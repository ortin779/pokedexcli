package pokeapi

import (
	"net/http"
	"time"

	"github.com/ortin779/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: *pokecache.NewCache(time.Minute),
	}
}
