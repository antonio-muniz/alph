package http

import "net/http"

type noContentHandler struct{}

func NewNoContentHandler() http.Handler {
	return noContentHandler{}
}

func (h noContentHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	httpResponse.WriteHeader(http.StatusNoContent)
}
