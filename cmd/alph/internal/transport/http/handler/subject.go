package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/sarulabs/di"
)

type createSubjectHandler struct {
	components di.Container
}

func NewCreateSubjectHandler(components di.Container) http.Handler {
	return createSubjectHandler{components: components}
}

func (h createSubjectHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	body, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}
	var request request.CreateSubject
	err = json.Unmarshal(body, &request)
	if err != nil {
		httpResponse.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := httpRequest.Context()
	err = controller.CreateSubject(ctx, h.components, request)
	switch err {
	case nil:
		httpResponse.WriteHeader(http.StatusCreated)
	default:
		fmt.Println(err.Error())
		httpResponse.WriteHeader(http.StatusInternalServerError)
		return
	}
}
