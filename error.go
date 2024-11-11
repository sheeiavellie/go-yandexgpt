package yandexgpt

type YandexGPTError struct {
	HTTPCode   int             `json:"httpCode"`
	Message    string          `json:"message"`
	HTTPStatus string          `json:"httpStatus"`
	Details    DetailsResponse `json:"error.details"`
}
