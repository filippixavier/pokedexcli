package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	Next       string
	Previous   string
	httpClient *http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		Next:     "",
		Previous: "",
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
