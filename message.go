package yandexgpt

type YandexGPTMessage struct {
	Role yandexGPTRole `json:"role"`
	Text string        `json:"text"`
}
