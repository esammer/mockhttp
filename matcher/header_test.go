package matcher

import (
	"testing"

	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
)

func TestHeaderMatcher_Match(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		matcher  Matcher
		expected bool
	}{
		{
			name:     "Empty header key",
			key:      "",
			value:    "test",
			matcher:  &HeaderMatcher{Header: "key1", Value: "test"},
			expected: false,
		},
		{
			name:     "header key match, value fail",
			key:      "key1",
			value:    "wrongValue",
			matcher:  &HeaderMatcher{Header: "key1", Value: "test"},
			expected: false,
		},
		{
			name:     "header key and value fail",
			key:      "key2",
			value:    "test2",
			matcher:  &HeaderMatcher{Header: "key1", Value: "test"},
			expected: false,
		},
		{
			name:     "header key and value match",
			key:      "key1",
			value:    "test",
			matcher:  &HeaderMatcher{Header: "key1", Value: "test"},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httputil.NewPostRequest("/", nil)

			req.Header.Set(test.key, test.value)

			require.Equal(t, test.expected, test.matcher.Match(req))
		})
	}
}
