package yandexgpt

import (
	"context"
	"fmt"
	"net/http"
)

const embeddingsURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/textEmbedding"

var (
	TextSearchDoc   = embeddingsModel{modelName: "text-search-doc/latest"}
	TextSearchQuery = embeddingsModel{modelName: "text-search-query/latest"}
)

type embeddingsModel struct {
	modelName string
}

func (em embeddingsModel) String() string {
	return fmt.Sprint(em.modelName)
}
func (c *YandexGPTClient) GetEmbedding(
	ctx context.Context,
	request YandexGPTEmbeddingsRequest,
) (response EmbeddingResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, embeddingsURL, request)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	return
}
