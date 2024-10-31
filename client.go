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
	config         *YandexGPTClientConfig
}

// Creates new YandexGPT Client.
//
// If you're using this option, keep in mind that you will need to generate IAM  token yourself.
func NewYandexGPTClientWithIAMToken(
	iamToken string,
) *YandexGPTClient {
	config := NewYandexGPTClientConfigWithIAMToken(iamToken)

	return &YandexGPTClient{
		config:         config,
		requestBuilder: internal.NewRequestBuilder(),
	}
}

// Creates new YandexGPT Client.
func NewYandexGPTClient() *YandexGPTClient {
	config := NewYandexGPTClientConfig()

	return &YandexGPTClient{
		config:         config,
		requestBuilder: internal.NewRequestBuilder(),
	}
}

// Creates new YandexGPT Client.
//
// You will need to specify your own API key.
func NewYandexGPTClientWithAPIKey(
	apiKey string,
) *YandexGPTClient {
	config := NewYandexGPTClientConfigWithAPIKey(apiKey)

	return &YandexGPTClient{
		config:         config,
		requestBuilder: internal.NewRequestBuilder(),
	}
}

// TODO: change type any maybe
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

	if c.config.IAMToken != "" {
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
