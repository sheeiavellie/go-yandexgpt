package yandexgpt

import "net/http"

const (
	completionURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
)

type YandexGPTClientConfig struct {
	ApiKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewYandexGPTClientConfig(
	apiKey string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		ApiKey:     apiKey,
		BaseURL:    completionURL,
		HTTPClient: &http.Client{},
	}
}
