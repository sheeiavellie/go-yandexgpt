package yandexgpt

type YandexGPTRequest struct {
	ModelURI          string                     `json:"modelUri"`
	CompletionOptions YandexGPTCompletionOptions `json:"completionOptions"`
	Messages          []YandexGPTMessage         `json:"messages"`
}

type YandexGPTCompletionOptions struct {
	Stream      bool    `json:"stream"`
	Temperature float32 `json:"temperature"`
	MaxTokens   int     `json:"maxTokens"`
}
