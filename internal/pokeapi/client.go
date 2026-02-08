package pokeapi

import (
	"github.com/vetal-bla/bootdev-pokedex/internal/pokecache"
	"net/http"
	"time"
)

// Client - struct
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient creation
func NewClient(timeout time.Duration) Client {
	globalCache := pokecache.NewCache(20 * time.Minute)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: globalCache,
	}
}
