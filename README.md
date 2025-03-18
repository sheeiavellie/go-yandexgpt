# ✨ Go YandexGPT ✨
[![Release CI](https://github.com/sheeiavellie/go-yandexgpt/actions/workflows/semantic-release.yml/badge.svg)](https://github.com/sheeiavellie/go-yandexgpt/actions/workflows/semantic-release.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/sheeiavellie/go-yandexgpt.svg)](https://pkg.go.dev/github.com/sheeiavellie/go-yandexgpt)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-openai)](https://goreportcard.com/report/github.com/sheeiavellie/go-yandexgpt)

This library provides unofficial Go client for [YandexGPT API](https://cloud.yandex.ru/en/services/yandexgpt).

## Installation

```
go get github.com/sheeiavellie/go-yandexgpt@latest
```
Currently, go-yandexgpt requires Go version 1.21 or greater.


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
				Text: "ONE",
			},
		},
	}

	response, err := client.GetCompletion(context.Background(), request)
	if err != nil {
		fmt.Println("Request error")
		return
	}

	fmt.Println(response.Result.Alternatives[0].Message.Text)
}

```

### Getting an API Key/IAM token:

You can get all the necessary information from the [official documentation](https://cloud.yandex.ru/en/docs/yandexgpt/quickstart).

## Contribution guideline

### Contributing
You can contribute by:
- Reporting issues
- Suggesting new features and enhancements
- Improving documentation

For **minor** changes you can just send a PR without opening linked issue.

For **major** changes open an issue.

### Commits and PRs
I highly encourage using [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/ "conventional commits") style in commit messages.

For PR titles it is *required* to use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/ "conventional commits") style titles.

You can use any of these prefixes:
- fix
- feat
- chore
- refactor
- test
- ci

## Credits 🖼️

- [Alexander V. Trotsky](https://github.com/sheeiavellie "Alexander V. Trotsky") (Author)
- [Contributors](https://github.com/sheeiavellie/go-yandexgpt/graphs/contributors "Contributors")

## Acknowledgement
#### Thank you very much

This project was ***highly*** inspired by [go-openai](https://github.com/sashabaranov/go-openai) by [sashabaranov](https://github.com/sashabaranov)
_______________
*Made with 💖 and some wizardry 🧙🔮*

