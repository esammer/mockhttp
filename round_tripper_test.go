package mockhttp

import (
	"fmt"
	"github.com/esammer/mockhttp/matcher"
	"github.com/esammer/mockhttp/responder"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewMockRoundTripper(t *testing.T) {
	rules := []*Rule{{}, {}}

	rt := NewMockRoundTripper(rules...)
	require.Len(t, rt.Rules, 2)
	require.Nil(t, rt.DefaultResponder)
}

func TestNewMockRoundTripperWithDefault(t *testing.T) {
	defResp := &responder.ConstantResponder{}
	rules := []*Rule{{}, {}}

	rt := NewMockRoundTripperWithDefault(defResp, rules...)
	require.NotNil(t, rt)
	require.Len(t, rt.Rules, 2)
	require.Equal(t, defResp, rt.DefaultResponder)
}

func TestMockRoundTripper_RoundTrip(t *testing.T) {
	rt := &MockRoundTripper{
		Rules: []*Rule{
			{
				Matcher: &matcher.AllOfMatcher{
					Matchers: []matcher.Matcher{
						&matcher.MethodMatcher{Method: http.MethodGet},
						&matcher.PathMatcher{Path: "/"},
						matcher.FuncMatcher(func(req *http.Request) bool {
							return req.URL.Query().Get("redirect") == "true"
						}),
					},
				},
				Responder: &responder.HandlerResponder{Handler: http.RedirectHandler("/redirected", http.StatusTemporaryRedirect)},
			},
			{
				Matcher: &matcher.AllOfMatcher{
					Matchers: []matcher.Matcher{
						&matcher.MethodMatcher{Method: http.MethodGet},
						&matcher.PathMatcher{Path: "/redirected"},
					},
				},
				Responder: &responder.ConstantResponder{StatusCode: http.StatusOK, Body: "redirected!"},
			},
			{
				Matcher: &matcher.AllOfMatcher{
					Matchers: []matcher.Matcher{
						&matcher.MethodMatcher{Method: http.MethodGet},
						&matcher.PathMatcher{Path: "/404"},
					},
				},
				Responder: &responder.ConstantResponder{StatusCode: http.StatusNotFound},
			},
			{
				Matcher: &matcher.AllOfMatcher{
					Matchers: []matcher.Matcher{
						&matcher.MethodMatcher{Method: http.MethodGet},
						&matcher.PathMatcher{Path: "/"},
					},
				},
				Responder: &responder.ConstantResponder{StatusCode: http.StatusOK},
			},
		},
	}

	client := &http.Client{
		Transport: rt,
	}

	resp, err := client.Get("/")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, int64(0), resp.ContentLength)
	b, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	require.Equal(t, "", string(b))

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Printf("Checking redirect to: %s from: %s\n", req.URL, via[0].URL)
		return nil
	}
	resp, err = client.Get("/?redirect=true")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, int64(11), resp.ContentLength)
	b, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	require.Equal(t, "redirected!", string(b))

	resp, err = client.Get("/?redirect=false")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, int64(0), resp.ContentLength)
	b, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	require.Equal(t, "", string(b))

	resp, err = client.Get("/404")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Equal(t, int64(0), resp.ContentLength)

	resp, err = client.Get("/no-match")
	require.Error(t, err)
	require.Nil(t, resp)
}

func TestMockRoundTripper_RoundTrip_DefaultResponder(t *testing.T) {
	rt := &MockRoundTripper{
		// We 202 Accept you no matter what! :)
		DefaultResponder: &responder.ConstantResponder{StatusCode: http.StatusAccepted},
	}

	client := &http.Client{
		Transport: rt,
	}

	resp, err := client.Get("/")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusAccepted, resp.StatusCode)
}

func ExampleMockRoundTripper_configuration() {
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

	// Requests that match a rule will produce the mock response.
	resp, err := client.Get("/")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
}
