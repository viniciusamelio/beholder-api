package presentation

import (
	"encoding/json"
	"io"
)

func ParseBody[S interface{}](b io.ReadCloser) (*S, error) {
	instance := new(S)
	bytes, err := io.ReadAll(b)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
