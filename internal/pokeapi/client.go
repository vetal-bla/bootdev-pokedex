package pokeapi

import (
	"net/http"
	"time"
)

// Client - struct
type Client struct {
	httpClient http.Client
}

// NewClient creation
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
