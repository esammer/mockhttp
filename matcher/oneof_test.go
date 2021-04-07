package matcher

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestOneOfMatcher_Match(t *testing.T) {
	tests := []struct {
		name     string
		matcher  OneOfMatcher
		expected bool
	}{
		{
			name:     "match empty",
			matcher:  OneOfMatcher{Matchers: []Matcher{}},
			expected: true,
		},
		{
			name: "match method",
			matcher: OneOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodGet},
			}},
			expected: true,
		},
		{
			name: "match path",
			matcher: OneOfMatcher{Matchers: []Matcher{
				&PathMatcher{Path: "/"},
			}},
			expected: true,
		},
		{
			name: "match first",
			matcher: OneOfMatcher{Matchers: []Matcher{
				&PathMatcher{Path: "/"},
				&MethodMatcher{Method: http.MethodPost},
			}},
			expected: true,
		},
		{
			name: "match second",
			matcher: OneOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodPost},
				&PathMatcher{Path: "/"},
			}},
			expected: true,
		},
		{
			name: "fail both",
			matcher: OneOfMatcher{Matchers: []Matcher{
				&MethodMatcher{Method: http.MethodPost},
				&PathMatcher{Path: "/404"},
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
