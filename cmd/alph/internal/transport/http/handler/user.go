package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/pkg/system"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
)

type createUserHandler struct {
	system system.System
}

func NewUserHandler(sys system.System) http.Handler {
	return createUserHandler{system: sys}
}

func (h createUserHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var request request.NewUser
	err := json.NewDecoder(httpRequest.Body).Decode(&request)
	if err != nil {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := httpRequest.Context()
	err = controller.CreateUser(ctx, h.system, request)
	switch err {
	case nil:
		httpResponse.WriteHeader(http.StatusCreated)
	default:
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}
}
