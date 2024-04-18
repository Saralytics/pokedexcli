package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	http.Client
}

func NewClient() Client {
	client := http.Client{}
	client.Timeout = time.Minute

	return Client{
		client,
	}
}
