package yandexgpt

import (
	"context"
	"net/http"
)

const operationURL = "https://operation.api.cloud.yandex.net/operations"

// Get operation status from yandex cloud. Use it if you
// are working with yandex api via async methods.
//
// If you're using IAM token, make sure to update client's IAM token by calling
// GetIAMToken(iamToken string) method first.
//
// Keep in mind that if for some strange reason you provided  API key and IAM token to the client,
// this method will use API key.
func (c *YandexGPTClient) GetOperationStatus(
	ctx context.Context,
	operationID string,
) (response OperationResponse, err error) {
	endpoint := operationURL + "/" + operationID

	req, err := c.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	return
}
