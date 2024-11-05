package yandexgpt

import (
	"context"
	"net/http"
)

const getIAMUrl = "https://iam.api.cloud.yandex.net/iam/v1/tokens"

// Updates IAM token.
//
// Always call it before creating a request.
//
// If you will use it when API key is specified, method CreateRequest(...) will always use API key.
func (c *YandexGPTClient) UpdateIAMToken(ctx context.Context) error {
	iamRq := YandexIAMRequest{APIKey: c.config.ApiKey}
	req, err := c.newRequest(ctx, http.MethodPost, getIAMUrl, iamRq)
	if err != nil {
		return err
	}

	var resp YandexIAMResponse
	err = c.sendRequest(req, &resp)
	if err != nil {
		return err
	}

	//set new IAMToken
	c.config.updateIAMToken(resp.IAMToken)
	return nil
}
