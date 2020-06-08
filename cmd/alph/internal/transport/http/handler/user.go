package handler

import (
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/middleware"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/pkg/system"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/respond"
)

type newUserHandler struct {
	system system.System
}

func NewUserHandler(sys system.System) http.Handler {
	return middleware.MessageParser(
		message.NewUserRequest{},
		newUserHandler{system: sys},
	)
}

func (h newUserHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	ctx := httpRequest.Context()
	request := ctx.Value(middleware.Message).(message.NewUserRequest)
	err := controller.NewUser(ctx, h.system, request)
	switch err {
	case nil:
		respond.Created(httpResponse, "<user-id>", message.NewUserResponse{})
	default:
		respond.InternalServerError(httpResponse)
	}
}
