package mockhttp

import (
	"errors"
	"fmt"
	"github.com/esammer/mockhttp/httputil"
	"github.com/esammer/mockhttp/matcher"
	"github.com/esammer/mockhttp/responder"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestNewRule(t *testing.T) {
	m := &matcher.MethodMatcher{Method: http.MethodGet}
	rs := &responder.ConstantResponder{}
	ru := NewRule(m, rs)

	require.NotNil(t, ru)
	require.Equal(t, m, ru.Matcher)
	require.Equal(t, rs, ru.Responder)

	require.PanicsWithError(t, "matcher can not be nil", func() {
		NewRule(nil, rs)
	})

	require.PanicsWithError(t, "responder can not be nil", func() {
		NewRule(m, nil)
	})
}

func TestRule_Match(t *testing.T) {
	rule := &Rule{
		Matcher: &matcher.PathMatcher{Path: "/"},
	}

	require.True(t, rule.Match(httputil.NewGetRequest("/")))
	require.False(t, rule.Match(httputil.NewGetRequest("/no-match")))
}

func TestRule_MatchAndInvoke(t *testing.T) {
	rule := &Rule{
		Matcher: &matcher.PathMatcher{Path: "/"},
		Responder: &responder.ConstantResponder{
			StatusCode: http.StatusOK,
		},
	}

	resp, ok := rule.MatchAndInvoke(httputil.NewGetRequest("/"))
	require.True(t, ok)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	resp, ok = rule.MatchAndInvoke(httputil.NewGetRequest("/no-match"))
	require.False(t, ok)
	require.Nil(t, resp)

	t.Run("panic on err", func(t *testing.T) {
		rule := &Rule{
			Matcher: &matcher.PathMatcher{Path: "/"},
			Responder: responder.FuncResponder(func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("boom")
			}),
		}

		require.PanicsWithError(t, "boom", func() {
			rule.MatchAndInvoke(httputil.NewGetRequest("/"))
		})
	})
}

// Most of the time you'll use create rules and configure MockRoundTripper to use them rather
// than invoking Match*() methods yourself, but if you're curious, here's how they work.
func ExampleRule_creatingRules() {
	// Rules combine a matcher with a responder.
	rule := NewRule(
		MatchPath("/"),
		RespondWithStatus(http.StatusOK),
	)

	// Let's pretend we have an http.Request.
	var req *http.Request

	// You can access the Matcher directly.
	if rule.Matcher.Match(req) {
		// req's path is /.
	}

	// But Rules are actually matchers themselves, and just delegate to the underlying
	// matcher. This is equivalent to the example above.
	if rule.Match(req) {
		// Same result!
	}

	// If you want to match and immediate invoke the responder, there's some syntactic
	// sugar for that as well. MatchAndInvoke() calls the responder when there's a match
	// and returns true or false to indicate the match status. Use it like you would
	// retrieve a value from a map for a key that might not exist.
	if resp, matched := rule.MatchAndInvoke(req); matched {
		// rule matched and resp will be non-nil.
		fmt.Printf("Received resp: %+v\n", resp)
	} else {
		fmt.Println("No match!")
	}
}
