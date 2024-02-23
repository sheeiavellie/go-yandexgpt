package yandexgpt

type YandexGPTClientConfig struct {
	apiKey string
}

func NewYandexGPTClientConfig(
	apiKey string,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		apiKey: apiKey,
	}
}
