package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/pkg/system"
)

type passwordAuthHandler struct {
	system system.System
}

func PasswordAuthHandler(sys system.System) http.Handler {
	return passwordAuthHandler{system: sys}
}

func (h passwordAuthHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var request request.PasswordAuth
	err := json.NewDecoder(httpRequest.Body).Decode(&request)
	if err != nil {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := httpRequest.Context()
	response, err := controller.PasswordAuth(ctx, h.system, request)
	switch err {
	case nil:
		responseBody, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err.Error())
			httpResponse.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = httpResponse.Write(responseBody)
		if err != nil {
			fmt.Println(err.Error())
			httpResponse.WriteHeader(http.StatusInternalServerError)
			return
		}
	case controller.ErrIncorrectCredentials:
		httpResponse.WriteHeader(http.StatusForbidden)
	default:
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}
}
