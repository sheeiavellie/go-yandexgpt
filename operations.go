package yandexgpt

import (
	"context"
	"net/http"
)

const operationURL = "https://operation.api.cloud.yandex.net/operations"

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
