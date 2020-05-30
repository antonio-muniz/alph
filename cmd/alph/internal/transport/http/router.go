package http

import (
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/handler"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/middleware"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

func Router(components di.Container) http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.PanicRecovery)
	router.Use(middleware.ContentNegotiation)

	router.
		Handle("/api/auth", handler.NewAuthenticateHandler(components)).
		Methods(http.MethodPost)

	router.
		Handle("/api/subjects", handler.NewCreateSubjectHandler(components)).
		Methods(http.MethodPost)

	return router
}
