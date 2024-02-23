package data

import "fmt"

// TODO: Create library for working with YandexGPT as independent project
// The best way to define request for YandexGPT is to create library
// It should be in a form of YandexGPT client, where underneath the hood
// it will just be standart http client with a convenient API
// To make it clear, all of the following content really deserves it's own lib
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

type YandexGPTMessage struct {
	Role yandexGPTRole `json:"role"`
	Text string        `json:"text"`
}

// Roles
var (
	YandexGPTMessageRoleSystem    = yandexGPTRole{role: "system"}
	YandexGPTMessageRoleUser      = yandexGPTRole{role: "user"}
	YandexGPTMessageRoleAssistant = yandexGPTRole{role: "assistant"}
)

type yandexGPTRole struct {
	role string
}

func (yr yandexGPTRole) String() string {
	return fmt.Sprint(yr.role)
}
