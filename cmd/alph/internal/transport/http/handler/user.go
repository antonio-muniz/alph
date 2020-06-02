package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/antonio-muniz/alph/pkg/validator"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
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
		httpResponse.WriteHeader(http.StatusBadRequest)
		_, err := httpResponse.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		return
	}
	validator := validator.New(validator.ErrorFieldFromJSONTag())
	validationResult, err := validator.Validate(request)
	if err != nil {
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		_, err := httpResponse.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		return
	}
	if validationResult.Invalid() {
		responseBody, err := json.Marshal(validationResult)
		if err != nil {
			fmt.Println(err.Error())
			httpResponse.WriteHeader(http.StatusInternalServerError)
			_, err := httpResponse.Write([]byte("{}"))
			if err != nil {
				panic(err)
			}
			return
		}
		httpResponse.WriteHeader(http.StatusBadRequest)
		_, err = httpResponse.Write(responseBody)
		if err != nil {
			panic(err)
		}
		return
	}
	ctx := httpRequest.Context()
	err = controller.CreateUser(ctx, h.system, request)
	switch err {
	case nil:
		httpResponse.WriteHeader(http.StatusCreated)
		_, err := httpResponse.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		_, err := httpResponse.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		return
	}
}
