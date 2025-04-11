package yandexgpt

import (
	"context"
	"errors"
	"net/http"
)

const getIAMUrl = "https://iam.api.cloud.yandex.net/iam/v1/tokens"

// A way of updating IAM token. Safe for concurrent use.
//
// It uses your OAuth to get IAM token.
//
// Always update IAM before a request.
func (c *YandexGPTClient) GetIAMTokenOAuth(ctx context.Context, oauthToken string) error {
	if c.config.Mode() != CfgModeIAMToken {
		return errors.New("can't get IAM token: ApiKey is used")
	}

	iamRq := YandexIAMRequest{OAuthToken: oauthToken}

	req, err := c.newRequest(ctx, http.MethodPost, getIAMUrl, iamRq)
	if err != nil {
		return err
	}

	var resp YandexIAMResponse
	err = c.sendRequest(req, &resp)
	if err != nil {
		return err
	}

	c.config.setIAMToken(resp.IAMToken)

	return nil
}

// A way of updating IAM token. Safe for concurrent use.
//
// It is a setter, that provides you with a way of updating IAM externaly.
//
// Always update IAM before a request.
func (c *YandexGPTClient) UpdateIAMToken(ctx context.Context, iamToken string) error {
	if c.config.Mode() != CfgModeIAMToken {
		return errors.New("can't update IAM token: ApiKey is used")
	}

	c.config.setIAMToken(iamToken)

	return nil
}

// A way of updating API key. Safe for concurrent use.
//
// It is a setter for API key.
func (c *YandexGPTClient) UpdateAPIKey(ctx context.Context, apiKey string) error {
	if c.config.Mode() != CfgModeApiKey {
		return errors.New("can't update API key: IAM token is used")
	}

	c.config.setAPIKey(apiKey)

	return nil
}
