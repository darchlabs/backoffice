package client

import "net/http"

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	client  HTTPClient
	baseURL string
}

type Config struct {
	Client  HTTPClient
	BaseURL string
}

func New(conf *Config) *Client {
	return &Client{
		client:  conf.Client,
		baseURL: conf.BaseURL,
	}
}
