package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/antonio-muniz/alph/pkg/encryption"
	"github.com/antonio-muniz/alph/pkg/jwt"

	"github.com/antonio-muniz/alph/pkg/models/token"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/encryption"
)

type AuthRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func main() {
	http.HandleFunc("/api/auth", func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				fmt.Println(err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			var authRequest AuthRequest
			err = json.Unmarshal(body, &authRequest)
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				return
			}
			if authRequest.Identity != "someone@example.org" || authRequest.Password != "123456" {
				response.WriteHeader(http.StatusForbidden)
				return
			}
			now := time.Now()
			token := token.Token{
				Header: token.Header{
					SignatureAlgorithm: "HS256",
					TokenType:          "JWT",
				},
				Payload: token.Payload{
					Audience:       "example.org",
					ExpirationTime: token.Timestamp(now.Add(30 * time.Minute)),
					IssuedAt:       token.Timestamp(now),
					Issuer:         "alph",
					Subject:        authRequest.Identity,
				},
			}
			encodedToken, err := jwt.Serialize(token)
			if err != nil {
				fmt.Println(err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			signedToken, err := jwt.Sign(encodedToken, "HS256", "zLcwW6w2MEwS8RMzP71azVbQJyOK4fiV")
			if err != nil {
				fmt.Println(err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			encryptedToken, err := encryption.AESEncrypt(
				signedToken,
				"dont-share-this-key-with-anybody",
			)
			if err != nil {
				fmt.Println(err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			encryptedAESKey, err := encryption.RSAEncrypt(
				"dont-share-this-key-with-anybody",
				fixtures.PublicKey(),
			)
			if err != nil {
				fmt.Println(err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			accessToken := strings.Join([]string{encryptedToken, encryptedAESKey}, ".")

			authResponse := AuthResponse{AccessToken: accessToken}
			responseBody, err := json.Marshal(authResponse)
			if err != nil {
				fmt.Println(err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = response.Write(responseBody)
			if err != nil {
				fmt.Println(err.Error())
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
