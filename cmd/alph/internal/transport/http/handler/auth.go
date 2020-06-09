package handler

import (
	"net/http"

	"github.com/antonio-muniz/alph/pkg/middleware"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/respond"
	"github.com/antonio-muniz/alph/pkg/system"
)

type passwordAuthHandler struct {
	system system.System
}

func PasswordAuthHandler(sys system.System) http.Handler {
	return middleware.MessageParser(
		message.PasswordAuthRequest{},
		passwordAuthHandler{system: sys},
	)
}

func (h passwordAuthHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	ctx := httpRequest.Context()
	request := ctx.Value(middleware.Message).(message.PasswordAuthRequest)
	response, err := controller.PasswordAuth(ctx, h.system, request)
	switch err {
	case nil:
		respond.OK(httpResponse, response)
	case controller.ErrIncorrectCredentials:
		respond.Forbidden(httpResponse)
	default:
		respond.InternalServerError(httpResponse)
	}
}
