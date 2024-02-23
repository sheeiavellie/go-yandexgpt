package yandexgpt

import "fmt"

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
