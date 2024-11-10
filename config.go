package yandexgpt

import "net/http"

type YandexGPTClientConfig struct {
	OAuthToken string
	ApiKey     string
	IAMToken   string
	HTTPClient *http.Client
}

func NewYandexGPTClientConfig() *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		HTTPClient: &http.Client{},
	}
}

func NewYandexGPTClientConfigWithIAMToken(
	iamToken string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		IAMToken:   iamToken,
		HTTPClient: &http.Client{},
	}
}

func NewYandexGPTClientConfigWithAPIKey(
	apiKey string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		ApiKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

func NewYandexGPTClientConfigWithOAuthToken(
	oauthToken string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		OAuthToken: oauthToken,
		HTTPClient: &http.Client{},
	}
}

// Setter for IAM token in config.
//
// Use it for manually updating token in config.
func (c *YandexGPTClientConfig) SetIAMToken(iamToken string) {
	c.IAMToken = iamToken
}
