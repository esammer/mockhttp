package matcher

import "net/http"

// A matcher that matches the HTTP method.
//
// In most cases, one of the <method>MethodMatcher singletons should be used.
type MethodMatcher struct {
	Method string
}

func (m *MethodMatcher) Match(req *http.Request) bool {
	return m.Method == req.Method
}
