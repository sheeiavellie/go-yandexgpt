package yandexgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sheeiavellie/go-yandexgpt/internal"
)

// YandexGPT Client
type YandexGPTClient struct {
	requestBuilder internal.RequestBuilder
	config         *yandexGPTClientConfig
}

// Creates new YandexGPT Client.
func New(cfg *yandexGPTClientConfig) *YandexGPTClient {
	return &YandexGPTClient{
		config:         cfg,
		requestBuilder: internal.NewRequestBuilder(),
	}
}

func (c *YandexGPTClient) newRequest(
	ctx context.Context,
	method,
	url string,
	req any,
) (*http.Request, error) {
	request, err := c.requestBuilder.Build(ctx, method, url, req)
	if err != nil {
		return nil, err
	}

	c.setHeaders(request)

	return request, nil
}

func (c *YandexGPTClient) sendRequest(
	request *http.Request,
	v Response,
) error {
	response, err := c.config.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if internal.IsBadStatusCode(response.StatusCode) {
		return c.handleResponseError(response)
	}

	if v != nil {
		v.SetHeader(response.Header)
	}

	return json.NewDecoder(response.Body).Decode(v)
}

func (c *YandexGPTClient) setHeaders(request *http.Request) {
	request.Header.Set("Content-Type", "application/json")

	switch c.config.Mode() {
	case CfgModeIAMToken:
		request.Header.Set(
			"Authorization",
			fmt.Sprintf("Bearer %s", c.config.IAMToken),
		)
	case CfgModeApiKey:
		request.Header.Set(
			"Authorization",
			fmt.Sprintf("Api-Key %s", c.config.ApiKey),
		)
	}
}

func (c *YandexGPTClient) handleResponseError(response *http.Response) error {
	var errResponse YandexGPTResponseBad
	err := json.NewDecoder(response.Body).Decode(&errResponse)
	if err != nil {
		return err
	}
	return fmt.Errorf(
		"bad response. Http Status: %d %s. Message: %s",
		errResponse.Error.HTTPCode,
		errResponse.Error.HTTPStatus,
		errResponse.Error.Message,
	)
}
