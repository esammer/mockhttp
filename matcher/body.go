package matcher

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// A matcher that matches the HTTP body and content type.
//
// Both ContentType and Body are optional. When specified, they must match. Note that the body will be read, closed,
// and replaced with a copy as a side effect of calling Match() causing this to be more expensive than many of the other
// matchers. To reduce the likelihood of matches, it's good to combine this with method and path matchers in an
// AllOfMatcher.
type BodyMatcher struct {
	ContentType string
	Body        string
}

func (m *BodyMatcher) Match(req *http.Request) bool {
	if m.ContentType != req.Header.Get("Content-Type") {
		return false
	}

	// If the request body is nil, we match on an empty Body string.
	if req.Body == nil {
		return m.Body == ""
	}

	buf := &bytes.Buffer{}
	r := io.TeeReader(req.Body, buf)
	req.Body = ioutil.NopCloser(buf)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		// FIXME: While unlikely to fail, this leaves the request in a bad state.
		return false
	}

	// FIXME: Swallowing this error is bad.
	if err := req.Body.Close(); err != nil {
		return false
	}

	return string(b) == m.Body
}
