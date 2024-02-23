package yandexgpt

type YandexGPTClientConfig struct {
	apiKey    string
	catalogID string
	model     yandexGPTModel
}

func NewYandexGPTClientConfig(
	apiKey,
	catalogID string,
	model yandexGPTModel,
) *YandexGPTClientConfig {
	return &YandexGPTClientConfig{
		apiKey:    apiKey,
		catalogID: catalogID,
		model:     model,
	}
}
