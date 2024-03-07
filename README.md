# ✨ Go YandexGPT ✨
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-openai)](https://goreportcard.com/report/github.com/sheeiavellie/go-yandexgpt)

This library provides unofficial Go client for [YandexGPT API](https://cloud.yandex.ru/en/services/yandexgpt).

## Installation

```
go get github.com/sheeiavellie/go-yandexgpt
```
Currently, go-openai requires Go version 1.22 or greater.


## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/sheeiavellie/go-yandexgpt"
)

func main() {
	client := yandexgpt.NewYandexGPTClientWithAPIKey("apiKey")
	request := yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI("catalogID", yandexgpt.YandexGPTModelLite),
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Stream:      false,
			Temperature: 0.7,
			MaxTokens:   2000,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleSystem,
				Text: "Every time you get ONE you answer just TWO",
			},
			{
				Role: yandexgpt.YandexGPTMessageRoleUser,
				Text: "TWO",
			},
		},
	}

	response, err := client.CreateRequest(context.Background(), request)
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(response.Result.Alternatives[0].Message.Text)
}

```

### Getting an API Key/IAM token:

You can get all the necessary information from the [official documentation](https://cloud.yandex.ru/en/docs/yandexgpt/quickstart).

## Contributing

1. Create an issue
2. Fork the repository
3. Do your magik
4. Open a pull request (please, put the link to previously opened issue in it's title)

## Thank you very much

This project was ***highly*** inspired by [go-openai](https://github.com/sashabaranov/go-openai) by [sashabaranov](https://github.com/sashabaranov)

*Made with 💖 and some wizardry 🧙🔮*
