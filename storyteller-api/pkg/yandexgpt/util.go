package yandexgpt

import "fmt"

func MakeModelURI(catalogID, string, model yandexGPTModel) string {
	return fmt.Sprintf("gpt://%s/%s", catalogID, model)
}
