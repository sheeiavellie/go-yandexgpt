package yandexgpt

import "net/http"

type YandexGPTClientConfig struct {
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

func (c *YandexGPTClientConfig) updateIAMToken(iamToken string) {
	c.IAMToken = iamToken
}
