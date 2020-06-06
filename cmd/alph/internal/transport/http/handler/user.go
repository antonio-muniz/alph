package handler

import (
	"encoding/json"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/antonio-muniz/alph/pkg/validator"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/respond"
)

type createUserHandler struct {
	system system.System
}

func NewUserHandler(sys system.System) http.Handler {
	return createUserHandler{system: sys}
}

func (h createUserHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var request message.NewUserRequest
	err := json.NewDecoder(httpRequest.Body).Decode(&request)
	if err != nil {
		respond.MalformedRequest(httpResponse)
		return
	}
	validator := validator.New(validator.ErrorFieldFromJSONTag())
	validationResult, err := validator.Validate(request)
	if err != nil {
		respond.InternalServerError(httpResponse)
		return
	}
	if validationResult.Invalid() {
		respond.InvalidRequestParameters(httpResponse, validationResult)
		return
	}
	ctx := httpRequest.Context()
	err = controller.NewUser(ctx, h.system, request)
	switch err {
	case nil:
		respond.Created(httpResponse, "<user-id>", message.NewUserResponse{})
	default:
		respond.InternalServerError(httpResponse)
	}
}
