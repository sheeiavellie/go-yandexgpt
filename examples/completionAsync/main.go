package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sheeiavellie/go-yandexgpt"
)

func main() {
	client := yandexgpt.NewYandexGPTClientWithAPIKey(os.Getenv("YAAPI_KEY"))

	// get, update and set iam token
	ctx := context.Background()
	err := client.UpdateIAMToken(ctx)
	if err != nil {
		log.Fatal(err)
	}

	request := yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI(os.Getenv("CATALOG_ID"), yandexgpt.YandexGPT4ModelLite),
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Stream:      false,
			Temperature: 0.7,
			MaxTokens:   100,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role:         yandexgpt.YandexGPTMessageRoleSystem,
				Text:         "Every time you get ONE you answer just TWO",
			},
			{
				Role:         yandexgpt.YandexGPTMessageRoleUser,
				Text:         "ONE",
			},
		},
	}

	response, err := client.RunCompletionAsync(ctx, request)
	if err != nil {
		log.Fatal(err)
	}

	isCompleted := false
	for !isCompleted {
		status, err := client.GetOperationStatus(ctx, response.ID)
		if err != nil {
			log.Fatal(err)
		}

		if status.Done {
			isCompleted = true
      fmt.Println("\n Chat answer: \n")
			fmt.Println(status.Response.Alternatives[0].Message.Text)
		} else {
      time.Sleep(5 * time.Second)
    }
	}
}
