package yandexgpt

import (
	"context"
	"net/http"

	"github.com/rep-co/fablescope-backend/storyteller-api/pkg/yandexgpt/internal"
)

// TODO: Move const somewhere where it will make sense
const (
	BaseURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
)

type YandexGPTClient struct {
	requestBuilder internal.RequestBuilder
	config         *YandexGPTClientConfig
}

func NewYandexGPTClient(
	apiKey string,
) *YandexGPTClient {
	config := NewYandexGPTClientConfig(apiKey)

	return &YandexGPTClient{
		config: config,
	}
}

func (c *YandexGPTClient) newRequest(
	ctx context.Context,
	method,
	url string,
) (*http.Request, error) {
}

func (c *YandexGPTClient) CreateRequest(
	ctx context.Context,
	request YandexGPTRequest,
) (response YandexGPTResponse, err error) {

	//TODO:
	//1. Validate Request
	//2. Create Request via c.newRequest(...)
	//3. Send request via c.sendRequest(...)
	//P.s. Use pointers
	return
}
