package yandexgpt

import (
	"context"
	"net/http"
)

const completionURL = "https://llm.api.cloud.yandex.net/foundationModels/v1"

// Get completion from YandexGPT.
//
// If you're using IAM token, make sure to update client's IAM token by calling
// GetIAMToken(iamToken string) method first.
//
// Keep in mind that if for some strange reason you provided  API key and IAM token to the client,
// this method will use API key.
func (c *YandexGPTClient) GetCompletion(
	ctx context.Context,
	request YandexGPTRequest,
) (*YandexGPTResponse, error) {

	endpoint := completionURL + "/completion"

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, request)
	if err != nil {
		return nil, err
	}

	var resp YandexGPTResponse
	err = c.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Get completion from YandexGPT with async method.
//
// If you're using IAM token, make sure to update client's IAM token by calling
// GetIAMToken(iamToken string) method first.
//
// Keep in mind that if for some strange reason you provided  API key and IAM token to the client,
// this method will use API key.
func (c *YandexGPTClient) RunCompletionAsync(
	ctx context.Context,
	request YandexGPTRequest,
) (*YandexCompletionResponse, error) {
	endpoint := completionURL + "/completionAsync"

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, request)
	if err != nil {
		return nil, err
	}

	var resp YandexCompletionResponse
	err = c.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
