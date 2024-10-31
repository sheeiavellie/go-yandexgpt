package yandexgpt

import (
	"context"
	"net/http"
)

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
	// TODO:
	// 1. Validate Request

	endpoint := c.config.BaseURL + "/completion"

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, request)
	if err != nil {
		return
	}
	err = c.sendRequest(req, &response)

	return
}
