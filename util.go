package yandexgpt

import "fmt"

// Use this for creating model uri.
func MakeModelURI(catalogID string, model yandexGPTModel) string {
	return fmt.Sprintf("gpt://%s/%s", catalogID, model)
}

// Use this for creating model uri for embeddings.
func MakeEmbModelURI(catalogID string, model embeddingsModel) string {
	return fmt.Sprintf("emb://%s/%s", catalogID, model)
}
