package yandexgpt

import (
	"net/http"
)

// TODO: Move const somewhere where it will make sense
const (
	BaseURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
)

type YandexGPTClient struct {
	config     *YandexGPTClientConfig
	httpClient *http.Client
}

func NewYandexGPTClient(
	apiKey string,
) *YandexGPTClient {
	client := &http.Client{}
	config := NewYandexGPTClientConfig(apiKey)

	return &YandexGPTClient{
		config:     config,
		httpClient: client,
	}
}

func (c *YandexGPTClient) CreateRequest(
	request YandexGPTRequest,
) (response YandexGPTResponse, err error) {

	//TODO:
	//1. Validate Request
	//2. Create Request via c.newRequest(...)
	//3. Send request via c.sendRequest(...)
	//P.s. Use pointers
	return
}
