package matcher

import "net/http"

// A matcher that matches the path of the URL.
type PathMatcher struct {
	Path string
}

func (m *PathMatcher) Match(req *http.Request) bool {
	return m.Path == req.URL.Path
}
