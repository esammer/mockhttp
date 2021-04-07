package mockhttp

import (
	"github.com/esammer/mockhttp/matcher"
	"net/http"
	"regexp"
)

// Creates a new "all of" matcher.
//
// An "all of" matcher returns true if and only if all of its child matchers
// return true. If no child matchers are supplied, it returns false.
func MatchAllOf(matchers ...matcher.Matcher) *matcher.AllOfMatcher {
	return &matcher.AllOfMatcher{
		Matchers: matchers,
	}
}

// Creates a new "one of" matcher.
//
// A "one of" matcher returns true if any of its child matchers rturn true.
// If no child matchers are supplied, it returns false.
func MatchOneOf(matchers ...matcher.Matcher) *matcher.OneOfMatcher {
	return &matcher.OneOfMatcher{
		Matchers: matchers,
	}
}

var (
	// Matches CONNECT requests.
	ConnectMethodMatcher = MatchMethod(http.MethodConnect)

	// Matches DELETE requests.
	DeleteMethodMatcher = MatchMethod(http.MethodDelete)

	// Matches GET requests.
	GetMethodMatcher = MatchMethod(http.MethodGet)

	// Matches HEAD requests.
	HeadMethodMatcher = MatchMethod(http.MethodHead)

	// Matches OPTIONS requests.
	OptionsMethodMatcher = MatchMethod(http.MethodOptions)

	// Matches PATCH requests.
	PatchMethodMatcher = MatchMethod(http.MethodPatch)

	// Matches POST requests.
	PostMethodMatcher = MatchMethod(http.MethodPost)

	// Matches PUT requests.
	PutMethodMatcher = MatchMethod(http.MethodPut)

	// Matches TRACE requests.
	TraceMethodMatcher = MatchMethod(http.MethodTrace)
)

// Creates a new method matcher.
//
// In most cases, one of the <method>MethodMatcher singletons should be used.
func MatchMethod(method string) *matcher.MethodMatcher {
	return &matcher.MethodMatcher{Method: method}
}

// Creates a new URL path matcher.
//
// Path matchers match a request whose URL path is equal to the configured path.
// The path component of the URL is as documented in, and implemented by, url.URL
// in the standard library.
func MatchPath(path string) *matcher.PathMatcher {
	return &matcher.PathMatcher{Path: path}
}

// Creates a new URL path regex matcher.
//
// Regex path matchers match a request whose URL path matches the configured pattern.
// The path component of the URL is as documented in, and implemented by, url.URL
// in the standard library.
func MatchPathRegex(pattern string) *matcher.PathRegexMatcher {
	return &matcher.PathRegexMatcher{Pattern: regexp.MustCompile(pattern)}
}
