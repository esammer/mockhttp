package httputil

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest(http.MethodGet, "/", nil)
	require.NotNil(t, req)
	require.Equal(t, http.MethodGet, req.Method)
	require.Equal(t, "/", req.URL.Path)

	require.PanicsWithError(
		t,
		`net/http: invalid method " space is illegal"`,
		func() {
			NewRequest(" space is illegal", "", nil)
		},
	)
}

func TestNewGetRequest(t *testing.T) {
	req := NewGetRequest("/")
	require.NotNil(t, req)
	require.Equal(t, http.MethodGet, req.Method)
	require.Equal(t, "/", req.URL.Path)
}

func TestNewPostRequest(t *testing.T) {
	const bodyContent = `{ "message": "body" }`

	req := NewPostRequest("/", strings.NewReader(bodyContent))
	require.NotNil(t, req)
	require.Equal(t, http.MethodPost, req.Method)
	require.Equal(t, "/", req.URL.Path)

	body, err := ioutil.ReadAll(req.Body)
	require.NoError(t, err)
	require.Equal(t, []byte(bodyContent), body)
	require.NoError(t, req.Body.Close())
}
