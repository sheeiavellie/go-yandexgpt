package yandexgpt

import (
	"encoding/json"
	"fmt"
)

var (
	YandexGPTMessageRoleSystem    = yandexGPTRole{Role: "system"}
	YandexGPTMessageRoleUser      = yandexGPTRole{Role: "user"}
	YandexGPTMessageRoleAssistant = yandexGPTRole{Role: "assistant"}
)

type yandexGPTRole struct {
	Role string `json:"role"`
}

func (yr yandexGPTRole) String() string {
	return fmt.Sprint(yr.Role)
}

func (yr yandexGPTRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(yr.Role)
}

func (yr *yandexGPTRole) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &yr.Role); err != nil {
		return err
	}
	return nil
}
