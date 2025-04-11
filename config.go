package yandexgpt

import (
	"errors"
	"net/http"
	"sync"
)

const (
	CfgModeApiKey   string = "ApiKeyMode"
	CfgModeIAMToken string = "IAMTokenMode"
)

type yandexGPTClientConfig struct {
	// Auth creds
	ApiKey   string
	IAMToken string

	// Http clietn to make requests
	HTTPClient *http.Client

	Mode func() string

	mu *sync.Mutex
}

var (
	CfgApiKey = func(apiKey string) *yandexGPTClientConfig {
		return &yandexGPTClientConfig{
			mu:         &sync.Mutex{},
			ApiKey:     apiKey,
			HTTPClient: &http.Client{},
			Mode: func() string {
				return CfgModeApiKey
			},
		}
	}

	CfgIAMToken = &yandexGPTClientConfig{
		mu:         &sync.Mutex{},
		HTTPClient: &http.Client{},
		Mode: func() string {
			return CfgModeIAMToken
		},
	}
)

// TODO: reexport config and add proper validation
func (c *yandexGPTClientConfig) Validate() error {
	if c.ApiKey == "" || c.IAMToken == "" {
		return errors.New("no Auth credential provided")
	}

	if c.HTTPClient == nil {
		return errors.New("no HTTP client provided")
	}

	return nil
}

// Setter for IAM token in config.
func (c *yandexGPTClientConfig) setIAMToken(iamToken string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.IAMToken = iamToken
}

// Setter for API key in config.
func (c *yandexGPTClientConfig) setAPIKey(apiKey string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ApiKey = apiKey
}
