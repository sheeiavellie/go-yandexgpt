package yandexgpt

import "fmt"

// TODO: add other models
var (
	YandexGPTModel     = yandexGPTModel{modelName: "yandexgpt"}
	YandexGPTModelLite = yandexGPTModel{modelName: "yandexgpt-lite"}
)

type yandexGPTModel struct {
	modelName string
}

func (ym yandexGPTModel) String() string {
	return fmt.Sprint(ym.modelName)
}
