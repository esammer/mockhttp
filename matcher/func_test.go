package matcher

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestFuncMatcher_Match(t *testing.T) {
	tests := []struct {
		name     string
		matcher  FuncMatcher
		expected bool
	}{
		{
			name: "trivial true",
			matcher: func(req *http.Request) bool {
				return true
			},
			expected: true,
		},
		{
			name: "trivial false",
			matcher: func(req *http.Request) bool {
				return false
			},
			expected: false,
		},
		{
			name: "match path",
			matcher: func(req *http.Request) bool {
				return req.URL.Path == "/"
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httputil.NewRequest(http.MethodGet, "/", nil)
			require.Equal(t, test.expected, test.matcher.Match(req))
		})
	}
}
