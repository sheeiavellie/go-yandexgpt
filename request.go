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

type YandexIAMRequest struct {
	OAuthToken string `json:"yandexPassportOauthToken"`
}

type YandexGPTEmbeddingsRequest struct {
	ModelURI string `json:"modelUri"`
	Text     string `json:"text"`
}
