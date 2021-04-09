package matcher

import (
	"github.com/esammer/mockhttp/httputil"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strings"
	"testing"
)

func TestBodyMatcher_Match(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		contentType string
		body        string
		matcher     Matcher
		expected    bool
	}{
		{
			name:        "content type and body match",
			contentType: "application/json",
			body:        `{ "a": "b", "c": 1 }`,
			matcher:     &BodyMatcher{ContentType: "application/json", Body: `{ "a": "b", "c": 1 }`},
			expected:    true,
		},
		{
			name:     "body match",
			body:     `{ "a": "b", "c": 1 }`,
			matcher:  &BodyMatcher{Body: `{ "a": "b", "c": 1 }`},
			expected: true,
		},
		{
			name:        "content type match",
			contentType: "application/json",
			matcher:     &BodyMatcher{ContentType: "application/json"},
			expected:    true,
		},
		{
			name:        "content type matches, body fail",
			contentType: "application/json",
			body:        `{ "a": "b", "c": 2 }`,
			matcher:     &BodyMatcher{ContentType: "application/json", Body: `{ "a": "b", "c": 1 }`},
			expected:    false,
		},
		{
			name:        "content type fail, body match",
			contentType: "text/html",
			body:        `{ "a": "b", "c": 1 }`,
			matcher:     &BodyMatcher{ContentType: "application/json", Body: `{ "a": "b", "c": 1 }`},
			expected:    false,
		},
		{
			name:        "content type and body fail",
			contentType: "text/html",
			body:        `{ "a": "b", "c": 2 }`,
			matcher:     &BodyMatcher{ContentType: "application/json", Body: `{ "a": "b", "c": 1 }`},
			expected:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bodyReader := ioutil.NopCloser(strings.NewReader(test.body))
			req := httputil.NewPostRequest("/", bodyReader)

			if test.contentType != "" {
				req.Header.Set("Content-Type", test.contentType)
			}

			require.Equal(t, test.expected, test.matcher.Match(req))
			// We purposefully test twice to make sure the body reader is reset.
			require.Equal(t, test.expected, test.matcher.Match(req))
		})
	}
}
