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

func NewYandexGPTClient(
	apiKey string,
) *YandexGPTClient {
	config := NewYandexGPTClientConfig(apiKey)

	return &YandexGPTClient{
		config: config,
	}
}

type requestOptions struct {
	body   any
	header http.Header
}

type requestOption func(*requestOptions)

func withBody(body any) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func (c *YandexGPTClient) newRequest(
	ctx context.Context,
	method,
	url string,
	options ...requestOption,
) (*http.Request, error) {
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}

	for _, option := range options {
		option(args)
	}
	request, err := c.requestBuilder.Build(ctx, method, url, args.body, args.header)
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
	//TODO: Mb check here if headers where set
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
	request.Header.Set("Authorization", fmt.Sprintf("Api-Key %s", c.config.ApiKey))
}

func (c *YandexGPTClient) handleResponseError(response *http.Response) error {

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

	req, err := c.newRequest(ctx, http.MethodPost, c.config.BaseURL, withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	return
}
