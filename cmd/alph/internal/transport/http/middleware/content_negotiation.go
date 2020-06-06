package middleware

import (
	"net/http"

	"github.com/antonio-muniz/alph/pkg/respond"
)

type contentNegotiation struct {
	nextHandler http.Handler
}

func ContentNegotiation(nextHandler http.Handler) http.Handler {
	return contentNegotiation{nextHandler: nextHandler}
}

func (m contentNegotiation) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	httpResponse.Header().Set("Content-Type", "application/json")
	requestContentType := httpRequest.Header.Get("Content-Type")
	if requestContentType != "application/json" {
		respond.UnsupportedContentType(httpResponse)
		return
	}
	m.nextHandler.ServeHTTP(httpResponse, httpRequest)
}
