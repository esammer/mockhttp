package matcher

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestAllOfMatcher_Match(t *testing.T) {
	tests := []struct {
		name     string
		matcher  AllOfMatcher
		expected bool
	}{
		{
			name:     "match empty",
			matcher:  AllOfMatcher{Matchers: []Matcher{}},
			expected: true,
		},
		{
			name: "match method",
			matcher: AllOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodGet},
			}},
			expected: true,
		},
		{
			name: "match path",
			matcher: AllOfMatcher{Matchers: []Matcher{
				&PathMatcher{Path: "/"},
			}},
			expected: true,
		},
		{
			name: "match method and path",
			matcher: AllOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodGet},
				&PathMatcher{Path: "/"},
			}},
			expected: true,
		},
		{
			name: "fail method",
			matcher: AllOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodPost},
				&PathMatcher{Path: "/"},
			}},
			expected: false,
		},
		{
			name: "fail path",
			matcher: AllOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodGet},
				&PathMatcher{Path: "/a"},
			}},
			expected: false,
		},
		{
			name: "fail method and path",
			matcher: AllOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodPost},
				&PathMatcher{Path: "/a"},
			}},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httputil.NewRequest(http.MethodGet, "/", nil)
			require.Equal(t, test.expected, test.matcher.Match(req))
		})
	}
}
