package mockhttp

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRespondWithStatus(t *testing.T) {
	r := RespondWithStatus(http.StatusAccepted)

	resp, err := r.Response(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusAccepted, resp.StatusCode)
}

func TestRespondWithBody(t *testing.T) {
	r := RespondWithBody(http.StatusCreated, "application/json", "{ }")

	resp, err := r.Response(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	b, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "{ }", string(b))
}

func TestRespondWithHandler(t *testing.T) {
	r := RespondWithHandler(http.RedirectHandler("/", http.StatusTemporaryRedirect))

	require.NotNil(t, r)
	resp, err := r.Response(httputil.NewGetRequest("/"))
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
}
