package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type RequestBuilder interface {
	Build(
		ctx context.Context,
		method,
		url string,
		body any,
	) (*http.Request, error)
}

type HTTPRequestBuilder struct {
}

func NewRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{}
}

func (b *HTTPRequestBuilder) Build(
	ctx context.Context,
	method,
	url string,
	body any,
) (request *http.Request, err error) {
	var bodyReader io.Reader
	if body != nil {
		if v, ok := body.(io.Reader); ok {
			bodyReader = v
		} else {
			var bytes bytes.Buffer
			err = json.NewEncoder(&bytes).Encode(body)
			if err != nil {
				return
			}
			bodyReader = &bytes
		}
	}
	request, err = http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return
	}

	return
}
