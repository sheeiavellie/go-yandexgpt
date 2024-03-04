package yandexgpt

import "net/http"

const (
	completionURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
)

type YandexGPTClientConfig struct {
	ApiKey     string
	IAMToken   string
	BaseURL    string
	HTTPClient *http.Client
}

func NewYandexGPTClientConfig() *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		BaseURL:    completionURL,
		HTTPClient: &http.Client{},
	}
}

func NewYandexGPTClientConfigWithIAMToken(
	iamToken string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		IAMToken:   iamToken,
		BaseURL:    completionURL,
		HTTPClient: &http.Client{},
	}
}

func NewYandexGPTClientConfigWithAPIKey(
	apiKey string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		ApiKey:     apiKey,
		BaseURL:    completionURL,
		HTTPClient: &http.Client{},
	}
}

func (c *YandexGPTClientConfig) UpdateIAMToken(iamToken string) {
	c.IAMToken = iamToken
}
