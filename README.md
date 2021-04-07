# mockhttp

Mock HTTP is a library for intercepting and matching requests made by `http.Client` and returning mock responses.

Like most mocking libraries, this library provides a set of reusable "matchers" that match HTTP requests and respond in
a number of ways without involving an actual HTTP server. It's primarily useful for testing client applications as
opposed to handlers, and was originally developed while building a REST client SDK.

Unlike a lot of other libraries, mockhttp doesn't ever use global state, nor does it muck around with things like
`http.DefaultClient`. Instead, it works by providing a rule-evaluating `http.Transport` mocker (`MockRoundTripper`)
which intercepts and processes requests. This design allows for parallel test execution without worrying about global
state pollution or concurrency issues.

See the `MockRoundTripper` example for how to set up `http.Client` with a few rules. For information on the available
matchers, see the `Match*()` convenience methods, and the `matcher` package. Responders also have convenience methods in
the form `RespondWith*()` with implementations that live in the `responder` package.

## Getting Started

Add mockhttp to your project:

    go get -u http://github.com/esammer/mockhttp

Configure your `http.Client` with a `MockRoundTripper` transport:

    // Create an http.Client, but use MockRoundTripper as the transport.
    client := &http.Client{
        Transport: NewMockRoundTripper(
            // Add as many rules as you'd like. They're evaluated in order.
            // Rules contain a matcher and a responder.
            NewRule(MatchPath("/404"), RespondWithStatus(http.StatusNotFound)),
    
            NewRule(
                // Matchers can be combined into "all of" and "one of" expressions
                // for complex matching logic.
                MatchAllOf(
                    PostMethodMatcher,
                    MatchPath("/my-resource"),
                ),
                RespondWithBody(
                    http.StatusInternalServerError,
                    "application/json",
                    `{ "things": [ "a", "b", "c" ] }`,
                ),
            ),
        ),
	}

Make HTTP requests as you normally would. You'll get your mocked responses instead.

	// Requests that match a rule will produce the mock response.
	resp, err := client.Get("/")
	if err != nil {
		// ...
	}

    fmt.Printf("Response: %+v\n", resp)

## License

mockhttp is licensed under the Apache License 2.0.

See LICENSE for details.
