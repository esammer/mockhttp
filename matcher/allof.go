package matcher

import (
	"net/http"
)

// A matcher that requires all child matchers to match.
type AllOfMatcher struct {
	Matchers []Matcher
}

func (m *AllOfMatcher) Match(req *http.Request) bool {
	for _, matcher := range m.Matchers {
		if !matcher.Match(req) {
			return false
		}
	}

	return true
}
