package matcher

import "net/http"

// A matcher that requires one child matcher to match.
type OneOfMatcher struct {
	Matchers []Matcher
}

func (m *OneOfMatcher) Match(req *http.Request) bool {
	for _, matcher := range m.Matchers {
		if matcher.Match(req) {
			return true
		}
	}

	// Make trivial success true like AllOfMatcher.
	return len(m.Matchers) == 0
}
