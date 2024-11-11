package yandexgpt

import (
	"net/http"
	"time"
)

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
	Message        YandexGPTMessage     `json:"message"`
	ToolResultList YandexToolResultList `json:"toolResultList"`
	Status         string               `json:"status"`
}

type YandexGPTUsage struct {
	InputTokens      string `json:"inputTextTokens"`
	CompletionTokens string `json:"completionTokens"`
	TotalTokens      string `json:"totalTokens"`
}

type YandexIAMResponse struct {
	IAMToken  string    `json:"iamToken"`
	ExpiresAt time.Time `json:"expiresAt"`
	httpHeader
}

type YandexCompletionResponse struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	CreatedAt   string            `json:"createdAt"`
	ModifiedAt  string            `json:"modifiedAt"`
	Done        bool              `json:"done"`
	Metadata    *MetadataResponse `json:"metadata"`
	Error       *StatusResponse   `json:"error"`
	Response    string            `json:"response"`
	httpHeader
}

type StatusResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Details DetailsResponse `json:"details"`
}

type MetadataResponse struct {
	Type   string `json:"@type"`
	DiskID string `json:"diskId"`
}

type DetailsResponse struct {
	Type      string `json:"@type"`
	RequestID string `json:"requestId"`
}

type OperationResponse struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	CreatedAt   string            `json:"createdAt"`
	CreatedBy   string            `json:"createdBy"`
	ModifiedAt  string            `json:"modifiedAt"`
	Done        bool              `json:"done"`
	Metadata    *MetadataResponse `json:"metadata"`
	Error       *StatusResponse   `json:"error"`
	Response    *YandexResponse   `json:"response"`
	httpHeader
}

type EmbeddingResponse struct {
	Embedding    []float64 `json:"embedding"`
	NumTokens    string    `json:"numTokens"`
	ModelVersion string    `json:"modelVersion"`
	httpHeader
}

type YandexResponse struct {
	Type         string                 `json:"@type"`
	Alternatives []YandexGPTAlternative `json:"alternatives"`
}
