package yandexgpt

type YandexGPTResponse struct {
	Result YandexGPTResult `json:"result"`
}

type YandexGPTResult struct {
	Alternatives []YandexGPTAlternative `json:"alternatives"`
	Usage        YandexGPTUsage         `json:"usage"`
	ModelVersion string                 `json:"modelVersion"`
}

type YandexGPTAlternative struct {
	Message YandexGPTMessage `json:"message"`
	Status  string           `json:"status"`
}

type YandexGPTUsage struct {
	InputTokens      string `json:"inputTextTokens"`
	CompletionTokens string `json:"completionTokens"`
	TotalTokens      string `json:"totalTokens"`
}
