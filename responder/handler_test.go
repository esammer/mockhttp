package responder

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func TestHandlerResponder_Response(t *testing.T) {
	tests := []struct {
		name         string
		handler      http.Handler
		expectedResp func() *http.Response
	}{
		{
			name:    "redirection",
			handler: http.RedirectHandler("/new", http.StatusTemporaryRedirect),
			expectedResp: func() *http.Response {
				w := NewResponseWriter()
				w.Header().Set("Location", "/new")
				w.WriteHeader(http.StatusTemporaryRedirect)
				return w.GetResponse()
			},
		},
		{
			name: "trivial ok",
			handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
			}),
			expectedResp: func() *http.Response {
				w := NewResponseWriter()
				w.WriteHeader(http.StatusOK)
				return w.GetResponse()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responder := &HandlerResponder{Handler: test.handler}
			req := httputil.NewPostRequest("/", strings.NewReader(`{ "message": "hello world" }`))
			resp, err := responder.Response(req)
			require.NoError(t, err)
			require.Equal(t, test.expectedResp(), resp)
		})
	}
}
