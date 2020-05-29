package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/pkg/models/request"
	"github.com/sarulabs/di"
)

type authenticateHandler struct {
	components di.Container
}

func NewAuthenticateHandler(components di.Container) http.Handler {
	return authenticateHandler{components: components}
}

func (h authenticateHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	body, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}
	var request request.Authenticate
	err = json.Unmarshal(body, &request)
	if err != nil {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := httpRequest.Context()
	authResponse, err := controller.Authenticate(ctx, h.components, request)
	switch err {
	case nil:
		responseBody, err := json.Marshal(authResponse)
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
