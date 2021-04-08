package mockhttp

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/esammer/mockhttp/matcher"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func TestMatchAllOf(t *testing.T) {
	m := MatchAllOf(matcher.FuncMatcher(func(req *http.Request) bool {
		return true
	}))

	require.NotNil(t, m)
	require.True(t, m.Match(httputil.NewGetRequest("/")))
}

func TestMatchOneOf(t *testing.T) {
	m := MatchOneOf(matcher.FuncMatcher(func(req *http.Request) bool {
		return true
	}))

	require.NotNil(t, m)
	require.True(t, m.Match(httputil.NewGetRequest("/")))
}

func TestMethodMatcher_Singletons(t *testing.T) {
	tests := []struct {
		matcher *matcher.MethodMatcher
		method  string
	}{
		{matcher: ConnectMethodMatcher, method: http.MethodConnect},
		{matcher: DeleteMethodMatcher, method: http.MethodDelete},
		{matcher: HeadMethodMatcher, method: http.MethodHead},
		{matcher: GetMethodMatcher, method: http.MethodGet},
		{matcher: OptionsMethodMatcher, method: http.MethodOptions},
		{matcher: PatchMethodMatcher, method: http.MethodPatch},
		{matcher: PostMethodMatcher, method: http.MethodPost},
		{matcher: PutMethodMatcher, method: http.MethodPut},
		{matcher: TraceMethodMatcher, method: http.MethodTrace},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			require.True(t, test.matcher.Match(httputil.NewRequest(test.method, "/", nil)))
		})
	}
}

func TestMatchMethod(t *testing.T) {
	m := MatchMethod(http.MethodPost)
	require.NotNil(t, m)
	require.True(t, m.Match(httputil.NewPostRequest("/", nil)))
	require.False(t, m.Match(httputil.NewGetRequest("/")))
}

func TestMatchPath(t *testing.T) {
	m := MatchPath("/")
	require.NotNil(t, m)
	require.True(t, m.Match(httputil.NewGetRequest("/")))
	require.False(t, m.Match(httputil.NewGetRequest("/not-found")))
}

func TestMatchPathRegex(t *testing.T) {
	m := MatchPathRegex("/[a-z]+$")
	require.NotNil(t, m)
	require.True(t, m.Match(httputil.NewGetRequest("/abc")))
	require.False(t, m.Match(httputil.NewGetRequest("/abc/")))
}

func TestMatchBody(t *testing.T) {
	m := MatchBody("application/json", "{ }")
	require.NotNil(t, m)

	r1 := httputil.NewPostRequest("/", strings.NewReader("{ }"))
	r1.Header.Set("Content-Type", "application/json")
	require.True(t, m.Match(r1))

	require.False(t, m.Match(httputil.NewPostRequest("/", nil)))
}
