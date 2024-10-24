package yandexgpt

import "fmt"

var (
	YandexGPT3Model     = yandexGPTModel{modelName: "yandexgpt"}
	YandexGPT3ModelLite = yandexGPTModel{modelName: "yandexgpt-lite"}
	YandexGPT4Model     = yandexGPTModel{modelName: "yandexgpt/rc"}
	YandexGPT4ModelLite = yandexGPTModel{modelName: "yandexgpt-lite/rc"}
	YandexGPT4Model32k  = yandexGPTModel{modelName: "yandexgpt-32k/rc"}
)

type yandexGPTModel struct {
	modelName string
}

func (ym yandexGPTModel) String() string {
	return fmt.Sprint(ym.modelName)
}
