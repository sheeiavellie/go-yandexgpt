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
	err := client.GetIAMTokenOAuth(ctx, "OAUTH_TOKE")
	if err != nil {
		panic(err)
	}

	request := yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI("CATALOG_ID", yandexgpt.YandexGPT4ModelLite),
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Stream:      false,
			Temperature: 0.7,
			MaxTokens:   100,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleSystem,
				Text: "Every time you get ONE you answer just TWO",
			},
			{
				Role: yandexgpt.YandexGPTMessageRoleUser,
				Text: "ONE",
			},
		},
	}

	response, err := client.GetCompletion(ctx, request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Result.Alternatives[0].Message.Text)
}
