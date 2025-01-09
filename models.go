package yandexgpt

import "fmt"

var (
	// Yandex GPT Pro 3rd generation
	YandexGPTModel = yandexGPTModel{modelName: "yandexgpt"}
	// Yandex GPT Lite 3rd generation
	YandexGPTModelLite = yandexGPTModel{modelName: "yandexgpt-lite"}
	// Yandex GPT Pro 4th generation
	YandexGPT4Model = yandexGPTModel{modelName: "yandexgpt/rc"}
	// Yandex GPT Lite 4th generation
	YandexGPT4ModelLite = yandexGPTModel{modelName: "yandexgpt-lite/rc"}
	// Yandex GPT Pro 32k 4th generation
	YandexGPT4Model32k = yandexGPTModel{modelName: "yandexgpt-32k/rc"}
	// Llama Lite 3rd generation
	LLAMA3Lite = yandexGPTModel{modelName: "llama-lite/latest"}
	// Llama  3rd generation
	LLAMA3 = yandexGPTModel{modelName: "llama/latest"}
)

type yandexGPTModel struct {
	modelName string
}

func (ym yandexGPTModel) String() string {
	return fmt.Sprint(ym.modelName)
}
