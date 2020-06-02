package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	scenarios := []struct {
		description          string
		requestBody          map[string]interface{}
		expectedStatusCode   int
		expectedResponseBody map[string]interface{}
		userIsExpected       bool
	}{
		{
			description: "responds_created_for_creating_a_valid_user",
			requestBody: map[string]interface{}{
				"username": "new.user@example.org",
				"password": "hakunamatata",
			},
			expectedStatusCode:   nethttp.StatusCreated,
			expectedResponseBody: map[string]interface{}{},
			userIsExpected:       true,
		},
		{
			description:        "responds_bad_request_and_validation_errors_for_invalid_parameters",
			requestBody:        map[string]interface{}{},
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
			sys, err := internal.System()
			require.NoError(t, err)
			router := http.Router(sys)
			requestBody, err := json.Marshal(scenario.requestBody)
			require.NoError(t, err)
			requestBodyReader := bytes.NewReader(requestBody)
			request, err := nethttp.NewRequest(nethttp.MethodPost, "/api/users", requestBodyReader)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			var responseBody map[string]interface{}
			err = json.NewDecoder(response.Body).Decode(&responseBody)
			require.NoError(t, err)
			require.Equal(t, scenario.expectedResponseBody, responseBody)
			if scenario.userIsExpected {
				database := sys.Get("database").(storage.Database)
				user, err := database.GetUser(ctx, scenario.requestBody["username"].(string))
				require.NoError(t, err)
				passwordMatch, err := password.Validate(scenario.requestBody["password"].(string), user.HashedPassword)
				require.NoError(t, err)
				require.True(t, passwordMatch)
			}
		})
	}
}
