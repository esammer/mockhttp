package matcher

import "net/http"

// A matcher that uses an arbitrary function.
type FuncMatcher func(req *http.Request) bool

func (m FuncMatcher) Match(req *http.Request) bool {
	return m(req)
}
