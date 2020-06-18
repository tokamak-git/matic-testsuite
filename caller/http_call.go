package caller

import (
	"bytes"
	"errors"
	"net/http"
)

type HTTPCall struct {
	Endpoint    string      `json:"endPoint"`
	Method      string      `json:"method"`
	ContentType string      `json:"contentType"`
	Headers     http.Header `json:"headers"`
	Body        string      `json:"body"`
}

func (h HTTPCall) Call() (*http.Response, error) {
	switch h.Method {
	case http.MethodGet:
		return http.Get(h.Endpoint)
	case http.MethodPost:
		return http.Post(h.Endpoint, h.ContentType, bytes.NewBufferString(h.Body))
	}
	return nil, errors.New("Invalid method")
}
