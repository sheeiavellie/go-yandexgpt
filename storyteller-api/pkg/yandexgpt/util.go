package yandexgpt

import "fmt"

// Use this for creating model uri.
func MakeModelURI(catalogID string, model yandexGPTModel) string {
	return fmt.Sprintf("gpt://%s/%s", catalogID, model)
}
