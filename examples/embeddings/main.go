package main

import (
	"context"
	"fmt"

	"github.com/sheeiavellie/go-yandexgpt"
)

func main() {
	client := yandexgpt.New(yandexgpt.CfgIAMToken)

	// get, update and set iam token
	ctx := context.Background()
	err := client.GetIAMTokenOAuth(ctx, "OATH_TOKEN")
	if err != nil {
		panic(err)
	}

	request := yandexgpt.YandexGPTEmbeddingsRequest{
		ModelURI: yandexgpt.MakeEmbModelURI("CATALOG_ID", yandexgpt.TextSearchDoc),
		Text:     "dog is not a cat",
	}

	response, err := client.GetEmbedding(ctx, request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Embedding) // print result vector like a slice of float64
}
