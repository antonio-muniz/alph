package handler

import (
	"encoding/json"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/respond"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/antonio-muniz/alph/pkg/validator"
)

type passwordAuthHandler struct {
	system system.System
}

func PasswordAuthHandler(sys system.System) http.Handler {
	return passwordAuthHandler{system: sys}
}

func (h passwordAuthHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var request message.PasswordAuthRequest
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
