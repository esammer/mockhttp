package httputil

import (
	"io"
	"net/http"
)

// Create an HTTP GET request.
//
// This is a convenience method for, and is equivalent to:
//   req := NewRequest(http.MethodGet, path, nil)
func NewGetRequest(path string) *http.Request {
	return NewRequest(http.MethodGet, path, nil)
}

// Create an HTTP POST request.
//
// This is a convenience method for, and is equivalent to:
//   req := NewRequest(http.MethodPost, path, body)
func NewPostRequest(path string, body io.Reader) *http.Request {
	return NewRequest(http.MethodPost, path, body)
}

// Create an HTTP request.
//
// Unlike http.NewRequest(), this method panics if any of the parameters are invalid making it easier to use in tests.
func NewRequest(method string, path string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		panic(err)
	}

	return req
}
