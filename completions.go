package yandexgpt

import (
	"context"
	"net/http"
)

const completionURL = "https://llm.api.cloud.yandex.net/foundationModels/v1"

// Get completion from YandexGPT.
//
// If you're using IAM token, make sure to update client's IAM token by calling
// UpdateIAMToken(iamToken string) method first.
//
// Keep in mind that if for some strange reason you provided  API key and IAM token to the client,
// this method will use API key
func (c *YandexGPTClient) GetCompletion(
	ctx context.Context,
	request YandexGPTRequest,
) (response YandexGPTResponse, err error) {

	endpoint := completionURL + "/completion"

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, request)
	if err != nil {
		return
	}
	err = c.sendRequest(req, &response)

	return
}

func (c *YandexGPTClient) RunCompletionAsync(
	ctx context.Context,
	request YandexGPTRequest,
) (response YandexCompletionResponse, err error) {
	endpoint := completionURL + "/completionAsync"

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, request)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	return
}
