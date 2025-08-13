package yandexgpt

import "fmt"

// Use this for creating model uri.
func MakeModelURI(catalogID, model, version string) string {
	return fmt.Sprintf("gpt://%s/%s/%s", catalogID, model, version)
}

// Use this for creating fine-tuned model uri.
func MakeFineTunedModelURI(catalogID, model, version, tuningSuffix string) string {
	return fmt.Sprintf("gpt://%s/%s/%s@%s", catalogID, model, version, tuningSuffix)
}

// Use this for creating model uri for embeddings.
func MakeEmbModelURI(catalogID string, model embeddingsModel) string {
	return fmt.Sprintf("emb://%s/%s", catalogID, model)
}
