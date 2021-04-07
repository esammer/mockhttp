package matcher

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestMethodMatcher_Match(t *testing.T) {
	tests := []struct {
		matcher  MethodMatcher
		method   string
		expected bool
	}{
		{matcher: MethodMatcher{Method: http.MethodGet}, method: http.MethodGet, expected: true},
		{matcher: MethodMatcher{Method: http.MethodGet}, method: http.MethodPost, expected: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s_%s", test.matcher.Method, test.method), func(t *testing.T) {
			req, _ := http.NewRequest(test.method, "/", nil)
			require.Equal(t, test.expected, test.matcher.Match(req))
		})
	}
}
