package pokeapi

import (
	"net/http"
	"time"

	"github.com/filippixavier/pokedexcli/internal/pokecache"
)

type Client struct {
	Next       *string
	Previous   *string
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Second),
	}
}
