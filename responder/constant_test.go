package responder

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestConstantResponder_Response(t *testing.T) {
	tests := []struct {
		name         string
		resp         *ConstantResponder
		expectedResp func() *http.Response // http.Response needs complex construction :(
	}{
		{
			name: "status only",
			resp: &ConstantResponder{StatusCode: http.StatusOK},
			expectedResp: func() *http.Response {
				w := NewResponseWriter()
				w.WriteHeader(http.StatusOK)
				return w.GetResponse()
			},
		},
		{
			name: "no parameters",
			resp: &ConstantResponder{},
			expectedResp: func() *http.Response {
				w := NewResponseWriter()
				w.WriteHeader(http.StatusOK)
				return w.GetResponse()
			},
		},
		{
			name: "content type only",
			resp: &ConstantResponder{ContentType: "application/json"},
			expectedResp: func() *http.Response {
				w := NewResponseWriter()
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				return w.GetResponse()
			},
		},
		{
			name: "body only",
			resp: &ConstantResponder{Body: `{ "status": "ok" }`},
			expectedResp: func() *http.Response {
				w := NewResponseWriter()
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{ "status": "ok" }`))
				return w.GetResponse()
			},
		},
		{
			name: "status + content type + body",
			resp: &ConstantResponder{
				StatusCode:  http.StatusAccepted,
				ContentType: "application/json",
				Body:        `{ "status": "ok" }`,
			},
			expectedResp: func() *http.Response {
				w := NewResponseWriter()
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusAccepted)
				_, _ = w.Write([]byte(`{ "status": "ok" }`))
				return w.GetResponse()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httputil.NewGetRequest("/")
			expectedResp := test.expectedResp()

			resp, err := test.resp.Response(req)
			require.NoError(t, err) // Even though this returns an error, we should never get one.
			require.Equal(t, expectedResp, resp)
		})
	}
}
