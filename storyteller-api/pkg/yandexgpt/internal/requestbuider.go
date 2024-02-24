package internal

import (
	"context"
	"net/http"
)

type RequestBuilder interface {
	Build(
		ctx context.Context,
		method,
		url string,
		body any,
		header http.Header,
	) (*http.Request, error)
}

type HTTPRequestBuilder struct {
	encoder Encoder
}

func NewRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		encoder: &JSONEncoder{},
	}
}

func (b *HTTPRequestBuilder) Build(
	ctx context.Context,
	method,
	url string,
	body any,
	header http.Header,
) (*http.Request, error) {

}
