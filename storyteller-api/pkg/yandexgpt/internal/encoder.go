package internal

import (
	"bytes"
	"encoding/json"
)

type Encoder interface {
	Encode(v any) (*bytes.Buffer, error)
}

type JSONEncoder struct {
}

func (je *JSONEncoder) Encode(v any) (*bytes.Buffer, error) {
	var data bytes.Buffer
	if err := json.NewEncoder(&data).Encode(v); err != nil {
		return nil, err
	}
	return &data, nil
}
