package qweathersdkgo

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultBaseURL = "https://api.qweather.com/v7"
	DefaultTimeout = 10 * time.Second
)

// Client represents the QWeather API client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// ClientOption defines a function to configure a Client
type ClientOption func(*Client) error

// NewClient creates a new QWeather API client with the given options
func NewClient(apiKey string, options ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	client := &Client{
		baseURL: DefaultBaseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(url string) ClientOption {
	return func(c *Client) error {
		if url == "" {
			return errors.New("base URL cannot be empty")
		}
		c.baseURL = url
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		if httpClient == nil {
			return errors.New("HTTP client cannot be nil")
		}
		c.httpClient = httpClient
		return nil
	}
}

// WithTimeout sets a custom timeout for the HTTP client
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) error {
		if timeout <= 0 {
			return errors.New("timeout must be positive")
		}
		c.httpClient.Timeout = timeout
		return nil
	}
}

// buildURL is a helper method to build the full URL for an API request
func (c *Client) buildURL(endpoint string, params map[string]string) (string, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("key", c.apiKey)
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
