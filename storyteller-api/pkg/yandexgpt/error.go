package yandexgpt

type YandexGPTError struct {
	GRPCCode   int    `json:"grpcCode"`
	HTTPCode   int    `json:"httpCode"`
	Message    string `json:"message"`
	HTTPStatus string `json:"httpStatus"`
}
