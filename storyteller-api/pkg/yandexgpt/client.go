package yandexgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rep-co/fablescope-backend/storyteller-api/pkg/yandexgpt/internal"
)

type YandexGPTClient struct {
	requestBuilder internal.RequestBuilder
	config         *YandexGPTClientConfig
}

func NewYandexGPTClient() *YandexGPTClient {
	config := NewYandexGPTClientConfig()

	return &YandexGPTClient{
		config:         config,
		requestBuilder: internal.NewRequestBuilder(),
	}
}

func NewYandexGPTClientWithAPIKey(
	apiKey string,
) *YandexGPTClient {
	config := NewYandexGPTClientConfigWithAPIKey(apiKey)

	return &YandexGPTClient{
		config:         config,
		requestBuilder: internal.NewRequestBuilder(),
	}
}

func (c *YandexGPTClient) newRequest(
	ctx context.Context,
	method,
	url string,
	req YandexGPTRequest,
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

	if c.config.ApiKey != "" {
		request.Header.Set(
			"Authorization",
			fmt.Sprintf("Api-Key %s", c.config.ApiKey),
		)
	} else if c.config.IAMToken != "" {
		request.Header.Set(
			"Authorization",
			fmt.Sprintf("Bearer %s", c.config.IAMToken),
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
		"bad response. Http Status %d %s message %s",
		errResponse.Error.HTTPCode,
		errResponse.Error.HTTPStatus,
		errResponse.Error.Message,
	)
}

func (c *YandexGPTClient) CreateRequest(
	ctx context.Context,
	request YandexGPTRequest,
) (response YandexGPTResponse, err error) {

	//TODO:
	//1. Validate Request

	req, err := c.newRequest(ctx, http.MethodPost, c.config.BaseURL, request)
	if err != nil {
		return
	}
	err = c.sendRequest(req, &response)

	return
}
