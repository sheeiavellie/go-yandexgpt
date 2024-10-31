package yandexgpt

import "net/http"

const (
	completionURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
  baseURL = "https://llm.api.cloud.yandex.net/foundationModels/v1"
)

type YandexGPTClientConfig struct {
	ApiKey     string
	IAMToken   string
	BaseURL    string
	HTTPClient *http.Client
}

func NewYandexGPTClientConfig() *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func NewYandexGPTClientConfigWithIAMToken(
	iamToken string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		IAMToken:   iamToken,
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func NewYandexGPTClientConfigWithAPIKey(
	apiKey string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		ApiKey:     apiKey,
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *YandexGPTClientConfig) updateIAMToken(iamToken string) {
	c.IAMToken = iamToken
}
