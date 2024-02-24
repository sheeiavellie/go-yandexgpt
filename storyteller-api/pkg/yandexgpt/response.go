package yandexgpt

import "net/http"

type Response interface {
	SetHeader(http.Header)
}

type httpHeader http.Header

func (h *httpHeader) SetHeader(header http.Header) {
	*h = httpHeader(header)
}

type YandexGPTResponse struct {
	Result YandexGPTResult `json:"result"`
	httpHeader
}

type YandexGPTResponseBad struct {
	Error YandexGPTError `json:"error"`
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
