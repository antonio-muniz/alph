package middleware

import (
	"net/http"

	"github.com/antonio-muniz/alph/pkg/respond"
)

type panicRecovery struct {
	nextHandler http.Handler
}

func PanicRecovery(nextHandler http.Handler) http.Handler {
	return panicRecovery{nextHandler: nextHandler}
}

func (m panicRecovery) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	defer recoverFromPanic(httpResponse)
	m.nextHandler.ServeHTTP(httpResponse, httpRequest)
}

func recoverFromPanic(httpResponse http.ResponseWriter) {
	panicValue := recover()
	if panicValue != nil {
		respond.InternalServerError(httpResponse)
	}
}
