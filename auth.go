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
// If you will use it when API key is specified, method GetCompletion(...) will always use API key.
func (c *YandexGPTClient) GetIAMToken(ctx context.Context) error {
	iamRq := YandexIAMRequest{OAuthToken: c.config.OAuthToken}
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
	c.config.SetIAMToken(resp.IAMToken)
	return nil
}
