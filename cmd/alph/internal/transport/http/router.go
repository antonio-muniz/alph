package http

import (
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/handler"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/middleware"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/gorilla/mux"
)

func Router(sys system.System) http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.PanicRecovery)
	router.Use(middleware.ContentNegotiation)

	router.
		Handle("/api/auth/password", handler.PasswordAuthHandler(sys)).
		Methods(http.MethodPost)

	router.
		Handle("/api/users", handler.NewUserHandler(sys)).
		Methods(http.MethodPost)

	return router
}
