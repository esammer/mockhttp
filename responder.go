package mockhttp

import (
	"github.com/esammer/mockhttp/responder"
	"net/http"
)

func RespondWithStatus(statusCode int) *responder.ConstantResponder {
	return &responder.ConstantResponder{StatusCode: statusCode}
}

func RespondWithBody(statusCode int, contentType string, body string) *responder.ConstantResponder {
	return &responder.ConstantResponder{
		StatusCode:  statusCode,
		ContentType: contentType,
		Body:        body,
	}
}

func RespondWithHandler(h http.Handler) *responder.HandlerResponder {
	return &responder.HandlerResponder{
		Handler: h,
	}
}
