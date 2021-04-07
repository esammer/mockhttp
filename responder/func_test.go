package responder

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestFuncResponder_Response(t *testing.T) {
	r := FuncResponder(func(req *http.Request) (*http.Response, error) {
		w := NewResponseWriter()
		return w.GetResponse(), nil
	})

	require.NotNil(t, r)
	resp, err := r.Response(httputil.NewGetRequest("/"))
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, http.StatusText(http.StatusOK), resp.Status)
	require.NotNil(t, resp.Body)
	require.True(t, resp.Close)
}
