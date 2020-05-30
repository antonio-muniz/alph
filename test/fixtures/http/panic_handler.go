package http

import "net/http"

type panicHandler struct{}

func NewPanicHandler() http.Handler {
	return panicHandler{}
}

func (h panicHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	panic("OH SNAP!")
}
