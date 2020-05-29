package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/pkg/models/request"
)

func main() {
	http.HandleFunc("/api/auth", func(httpResponse http.ResponseWriter, httpRequest *http.Request) {
		switch httpRequest.Method {
		case http.MethodPost:
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
			authResponse, err := controller.Authenticate(ctx, request)
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
		default:
			httpResponse.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
