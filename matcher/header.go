package matcher

import "net/http"

// A matcher that matches the HTTP headers.
//
// Both Key is mandatory. When specified, they must match. Note that the body will be read, closed,
// and replaced with a copy as a side effect of calling Match() causing this to be more expensive than many of the other
// matchers. To reduce the likelihood of matches, it's good to combine this with method and path matchers in an
// AllOfMatcher.

type HeaderMatcher struct {
	Header string
	Value  string
}

func (h *HeaderMatcher) Match(req *http.Request) bool {

	actualValue := req.Header.Get(h.Header)
	if actualValue == "" {
		return false
	}

	if actualValue != h.Value {
		return false
	}

	return true

}
