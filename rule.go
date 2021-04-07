package mockhttp

import (
	"errors"
	"github.com/esammer/mockhttp/matcher"
	"github.com/esammer/mockhttp/responder"
	"net/http"
)

// A request/response rule.
//
// A rule pairs a matcher with a responder. When the former matches a request, the latter response is invoked to
// generate a response. For convenience, you're encouraged to use NewRule() to create rules.
type Rule struct {
	Matcher   matcher.Matcher
	Responder responder.Responder
}

// Create a new rule.
//
// Panics if either matcher or responder are nil.
func NewRule(matcher matcher.Matcher, responder responder.Responder) *Rule {
	if matcher == nil {
		panic(errors.New("matcher can not be nil"))
	}

	if responder == nil {
		panic(errors.New("responder can not be nil"))
	}

	return &Rule{
		Matcher:   matcher,
		Responder: responder,
	}
}

// A convenience method that invokes the configured matcher.
//
// This method is equivalent to:
//   r.Matcher.Match(req)
func (r *Rule) Match(req *http.Request) bool {
	return r.Matcher.Match(req)
}

// A convenience method that invokes the responder if the matcher matches.
//
// If the configured matcher matches, the responder is invoked and its response is returned along with a true value
// indicating there was a match. If there is no match the response is nil and false is returned. Any error produced by
// the responder causes a panic.
func (r *Rule) MatchAndInvoke(req *http.Request) (*http.Response, bool) {
	if r.Matcher.Match(req) {
		resp, err := r.Responder.Response(req)
		if err != nil {
			panic(err)
		}

		return resp, true
	}

	return nil, false
}
