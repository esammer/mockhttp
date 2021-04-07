package responder

import (
	"net/http"
)

// A responder that produces a static response.
//
// All fields are optional. When StatusCode is omitted, the response will be a 200 OK.
type ConstantResponder struct {
	StatusCode  int
	ContentType string
	Body        string
}

func (r *ConstantResponder) Response(req *http.Request) (*http.Response, error) {
	w := NewResponseWriter()

	if r.ContentType != "" {
		w.Header().Set("Content-Type", r.ContentType)
	}

	w.WriteHeader(r.StatusCode)

	if r.Body != "" {
		// w.Write() never returns an error.
		_, _ = w.Write([]byte(r.Body))
	}

	return w.GetResponse(), nil
}
