package responder

import (
	"net/http"
)

// A responder adaptor for http.Handler.
//
// This responder allows for the use of a real http.Handler, including middleware.
type HandlerResponder struct {
	Handler http.Handler
}

func (r *HandlerResponder) Response(req *http.Request) (*http.Response, error) {
	w := NewResponseWriter()
	r.Handler.ServeHTTP(w, req)
	return w.GetResponse(), nil
}
