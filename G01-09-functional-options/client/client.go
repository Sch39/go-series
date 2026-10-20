package client

import (
	"errors"
	"time"
)

type Client struct {
	Endpoint string
	Timeout  time.Duration
	Retries  int
}

type Option func(*Client) error

func New(options ...Option) (*Client, error) {
	client := &Client{
		Endpoint: "https://api.example.com",
		Timeout:  5 * time.Second,
		Retries:  2,
	}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}
	return client, nil
}

func WithTimeout(timeout time.Duration) Option {
	return func(client *Client) error {
		if timeout <= 0 {
			return errors.New("timeout must be positive")
		}
		client.Timeout = timeout
		return nil
	}
}

func WithRetries(retries int) Option {
	return func(client *Client) error {
		if retries < 0 {
			return errors.New("retries must not be negative")
		}
		client.Retries = retries
		return nil
	}
}

func WithEndpoint(endpoint string) Option {
	return func(client *Client) error {
		if endpoint == "" {
			return errors.New("endpoint must not be empty")
		}
		client.Endpoint = endpoint
		return nil
	}
}
