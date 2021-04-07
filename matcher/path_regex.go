package matcher

import (
	"net/http"
	"regexp"
)

type PathRegexMatcher struct {
	Pattern *regexp.Regexp
}

func (m *PathRegexMatcher) Match(req *http.Request) bool {
	return m.Pattern.MatchString(req.URL.Path)
}
