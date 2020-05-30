package middleware

import (
	"net/http"
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
		httpResponse.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	m.nextHandler.ServeHTTP(httpResponse, httpRequest)
}
