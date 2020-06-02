package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
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
		httpResponse.WriteHeader(http.StatusBadRequest)
		httpResponse.Write([]byte("{}"))
		return
	}
	validator := validator.New(validator.ErrorFieldFromJSONTag())
	validationResult, err := validator.Validate(request)
	if err != nil {
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		httpResponse.Write([]byte("{}"))
		return
	}
	if validationResult.Invalid() {
		responseBody, err := json.Marshal(validationResult)
		if err != nil {
			fmt.Println(err.Error())
			httpResponse.WriteHeader(http.StatusInternalServerError)
			httpResponse.Write([]byte("{}"))
			return
		}
		httpResponse.WriteHeader(http.StatusBadRequest)
		httpResponse.Write(responseBody)
	}
	ctx := httpRequest.Context()
	response, err := controller.PasswordAuth(ctx, h.system, request)
	switch err {
	case nil:
		responseBody, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err.Error())
			httpResponse.WriteHeader(http.StatusInternalServerError)
			httpResponse.Write([]byte("{}"))
			return
		}
		_, err = httpResponse.Write(responseBody)
		if err != nil {
			fmt.Println(err.Error())
			httpResponse.WriteHeader(http.StatusInternalServerError)
			httpResponse.Write([]byte("{}"))
			return
		}
	case controller.ErrIncorrectCredentials:
		httpResponse.WriteHeader(http.StatusForbidden)
		httpResponse.Write([]byte("{}"))
	default:
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		httpResponse.Write([]byte("{}"))
		return
	}
}
