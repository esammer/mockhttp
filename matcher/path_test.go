package matcher

import (
	"fmt"
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPathMatcher_Match(t *testing.T) {
	tests := []struct {
		matcher  PathMatcher
		path     string
		expected bool
	}{
		{matcher: PathMatcher{Path: ""}, path: "", expected: true},
		{matcher: PathMatcher{Path: "/"}, path: "/", expected: true},
		{matcher: PathMatcher{Path: "/a"}, path: "/a", expected: true},
		{matcher: PathMatcher{Path: "/"}, path: "/a", expected: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s_%s", test.matcher.Path, test.path), func(t *testing.T) {
			req := httputil.NewGetRequest(test.path)
			require.Equal(t, test.expected, test.matcher.Match(req))
		})
	}
}
