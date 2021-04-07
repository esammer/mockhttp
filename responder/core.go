package responder

import "net/http"

type Responder interface {
	Response(req *http.Request) (*http.Response, error)
}
