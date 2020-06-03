package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
	"github.com/antonio-muniz/alph/pkg/password"
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
		accessTokenIsExpected bool
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
			expectedStatusCode:    nethttp.StatusOK,
			accessTokenIsExpected: true,
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
			sys, err := internal.System(ctx)
			require.NoError(t, err)
			hashedCorrectPassword, err := password.Hash(scenario.correctPassword)
			require.NoError(t, err)
			database := sys.Get("database").(storage.Database)
			user := model.User{
				Username:       scenario.correctUsername,
				HashedPassword: hashedCorrectPassword,
			}
			err = database.CreateUser(ctx, user)
			require.NoError(t, err)
			router := http.Router(sys)
			requestBody, err := json.Marshal(scenario.requestBody)
			require.NoError(t, err)
			requestBodyReader := bytes.NewReader(requestBody)
			request, err := nethttp.NewRequest(nethttp.MethodPost, "/api/auth/password", requestBodyReader)
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			var responseBody map[string]interface{}
			err = json.NewDecoder(response.Body).Decode(&responseBody)
			require.NoError(t, err)
			if scenario.accessTokenIsExpected {
				accessToken := responseBody["access_token"]
				require.IsType(t, "", accessToken)
				require.NotEmpty(t, accessToken)
			} else {
				require.Equal(t, scenario.expectedResponseBody, responseBody)
			}
		})
	}
}
