package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sheeiavellie/go-yandexgpt"
)

func main() {
	client := yandexgpt.New(yandexgpt.CfgIAMToken)

	// get, update and set iam token
	ctx := context.Background()
	err := client.GetIAMTokenOAuth(ctx, "OAUTH_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	request := yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI("CATALOG_ID", yandexgpt.YandexGPTLite, yandexgpt.VersionLatest),
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

	response, err := client.RunCompletionAsync(ctx, request)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation was cancelled")
			return

		case <-ticker.C:
			status, err := client.GetOperationStatus(ctx, response.ID)
			if err != nil {
				log.Fatal(err)
			}

			if status.Done {
				fmt.Println("Chat answer:")
				fmt.Println(status.Response.Alternatives[0].Message.Text)
				return
			}
		}
	}
}
