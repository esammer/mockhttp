package matcher

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestPathRegexMatcher_Match(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		expected bool
	}{
		{
			name:     "/.* on / matches",
			pattern:  `/.*`,
			expected: true,
		},
		{
			name:     "/[^/]+/[a-z]+ on / matches",
			pattern:  `/[^/]+/[a-z]+`,
			expected: true,
		},
		{
			name:     "/$ on / fails",
			pattern:  `/$`,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := &PathRegexMatcher{
				Pattern: regexp.MustCompile(test.pattern),
			}

			require.Equal(t, test.expected, m.Match(httputil.NewGetRequest("/a/sample/path")))
		})
	}
}
