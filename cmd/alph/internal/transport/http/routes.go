package http

import (
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/handler"
	"github.com/gorilla/mux"
)

func Router() http.Handler {
	router := mux.NewRouter()

	router.
		HandleFunc("/api/auth", handler.Authenticate).
		Methods(http.MethodPost)

	return router
}
