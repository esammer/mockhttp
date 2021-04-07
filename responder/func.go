package responder

import "net/http"

// A responder adapter that invokes a function.
type FuncResponder func(req *http.Request) (*http.Response, error)

func (r FuncResponder) Response(req *http.Request) (*http.Response, error) {
	return r(req)
}
