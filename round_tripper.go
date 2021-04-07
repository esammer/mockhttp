package mockhttp

import (
	"errors"
	"github.com/esammer/mockhttp/responder"
	"net/http"
)

// A request/response mocking http.Transport.
//
// The MockRoundTripper can be used in place of the default http.Transport to intercept requests and provide mocked
// responses based on a set of configured Rules, each of which contain a matcher.Matcher and responder.Responder pair.
// Optionally, a DefaultResponder can be configured which is used when no matching rule exists. When there is no default
// responder, failing to match a rule will produce an ErrNoMatchingRule error.
//
// Rules are evaluated in the order they are defined. RoundTrip() runs in O(n) time (where n = len(Rules)), but n is
// typically very small, so this shouldn't be an issue.
//
// See also Rule, matcher.Matcher, responder.Responder.
type MockRoundTripper struct {
	Rules            []*Rule
	DefaultResponder responder.Responder
}

var ErrNoMatchingRule = errors.New("no matching rule for request")

// Create a new MockRoundTripper with the given rules and a default responder.
func NewMockRoundTripperWithDefault(defaultResponder responder.Responder, rules ...*Rule) *MockRoundTripper {
	return &MockRoundTripper{
		DefaultResponder: defaultResponder,
		Rules:            rules,
	}
}

// Create a new MockRoundTripper with the given rules.
//
// See MockRoundTripper.
func NewMockRoundTripper(rules ...*Rule) *MockRoundTripper {
	return NewMockRoundTripperWithDefault(nil, rules...)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, rule := range m.Rules {
		if rule.Match(req) {
			return rule.Responder.Response(req)
		}
	}

	if m.DefaultResponder != nil {
		return m.DefaultResponder.Response(req)
	}

	return nil, ErrNoMatchingRule
}
