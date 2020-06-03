package handler_test

import (
	"context"
	nethttp "net/http"
	"testing"
	"time"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/test/internalhelpers"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/test/helpers"
	"github.com/stretchr/testify/require"
)

func TestPasswordAuth(t *testing.T) {
	scenarios := []struct {
		description           string
		correctUsername       string
		correctPassword       string
		correctClientID       string
		correctClientSecret   string
		requestBody           map[string]interface{}
		expectedStatusCode    int
		expectedUnpackedToken jwt.Token
		expectedResponseBody  map[string]interface{}
	}{
		{
			description:         "responds_ok_with_access_token_for_correct_credentials",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			requestBody: map[string]interface{}{
				"username":      "someone@example.org",
				"password":      "123456",
				"client_id":     "the-client",
				"client_secret": "the-client-is-scared-of-the-dark",
			},
			expectedStatusCode: nethttp.StatusOK,
			expectedUnpackedToken: jwt.Token{
				Issuer:         "alph",
				Audience:       "example.org",
				Subject:        "someone@example.org",
				ExpirationTime: jwt.Timestamp(helpers.Now().Add(30 * time.Minute)),
			},
		},
		{
			description:         "responds_forbidden_for_incorrect_credentials",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			requestBody: map[string]interface{}{
				"username":      "someone@example.org",
				"password":      "654321",
				"client_id":     "the-client",
				"client_secret": "the-client-is-scared-of-the-dark",
			},
			expectedStatusCode:   nethttp.StatusForbidden,
			expectedResponseBody: map[string]interface{}{},
		},
		{
			description:         "responds_bad_request_and_validation_errors_for_invalid_parameters",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			requestBody: map[string]interface{}{
				"client_id":     "the-client",
				"client_secret": "the-client-is-scared-of-the-dark",
			},
			expectedStatusCode: nethttp.StatusBadRequest,
			expectedResponseBody: map[string]interface{}{
				"validation_errors": []interface{}{
					map[string]interface{}{
						"type":  "MISSING",
						"field": "username",
					},
					map[string]interface{}{
						"type":  "MISSING",
						"field": "password",
					},
				},
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			ctx := context.Background()
			sys := internalhelpers.InitializeSystem(t, ctx)
			user := model.User{
				Username:       scenario.correctUsername,
				HashedPassword: helpers.HashPassword(t, scenario.correctPassword),
			}
			internalhelpers.CreateUser(t, ctx, sys, user)
			request := helpers.BuildHttpRequest(t,
				nethttp.MethodPost,
				"/api/auth/password",
				scenario.requestBody,
			)
			response := internalhelpers.ExecuteHttpRequest(t, sys, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			responseBody := helpers.DeserializeHttpResponseBody(t, response)
			if scenario.expectedStatusCode == nethttp.StatusOK {
				accessToken := responseBody["access_token"].(string)
				internalhelpers.VerifyAccessToken(t,
					sys,
					scenario.expectedUnpackedToken,
					accessToken,
				)
			} else {
				require.Equal(t, scenario.expectedResponseBody, responseBody)
			}
		})
	}
}
