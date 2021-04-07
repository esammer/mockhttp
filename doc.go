// Mock HTTP is a library for intercepting and matching requests made by `http.Client` and returning mock responses.
//
// Like most mocking libraries, this library provides a set of reusable "matchers" that match HTTP requests and respond in
// a number of ways without involving an actual HTTP server. It's primarily useful for testing client applications as
// opposed to handlers, and was originally developed while building a REST client SDK.
//
// Unlike a lot of other libraries, mockhttp doesn't ever use global state, nor does it muck around with things like
// `http.DefaultClient`. Instead, it works by providing a rule-evaluating `http.Transport` mocker (`MockRoundTripper`)
// which intercepts and processes requests. This design allows for parallel test execution without worrying about global
// state pollution or concurrency issues.
//
// See the `MockRoundTripper` example for how to set up `http.Client` with a few rules. For information on the available
// matchers, see the `Match*()` convenience methods, and the `matcher` package. Responders also have convenience methods in
// the form `RespondWith*()` with implementations that live in the `responder` package.
package mockhttp
