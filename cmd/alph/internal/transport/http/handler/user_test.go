package handler_test

import (
	"context"
	nethttp "net/http"
	"testing"

	"github.com/antonio-muniz/alph/test/helpers"

	"github.com/antonio-muniz/alph/cmd/alph/internal/test/internalhelpers"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	scenarios := []struct {
		description          string
		requestBody          map[string]interface{}
		expectedStatusCode   int
		expectedResponseBody map[string]interface{}
	}{
		{
			description: "responds_created_for_creating_a_valid_user",
			requestBody: map[string]interface{}{
				"username": "new.user@example.org",
				"password": "hakunamatata",
			},
			expectedStatusCode:   nethttp.StatusCreated,
			expectedResponseBody: map[string]interface{}{},
		},
		{
			description: "responds_bad_request_and_validation_errors_for_invalid_parameters",
			requestBody: map[string]interface{}{
				"username": "new.user@example.org",
			},
			expectedStatusCode: nethttp.StatusBadRequest,
			expectedResponseBody: map[string]interface{}{
				"validation_errors": []interface{}{
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
			request := helpers.BuildHttpRequest(t,
				nethttp.MethodPost,
				"/api/users",
				scenario.requestBody,
			)
			response := internalhelpers.ExecuteHttpRequest(t, sys, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			responseBody := helpers.DeserializeHttpResponseBody(t, response)
			require.Equal(t, scenario.expectedResponseBody, responseBody)
			expectedUsername := scenario.requestBody["username"].(string)
			if scenario.expectedStatusCode == nethttp.StatusCreated {
				expectedPassword := scenario.requestBody["password"].(string)
				internalhelpers.VerifyUserExists(t,
					ctx,
					sys,
					expectedUsername,
					expectedPassword,
				)
			} else {
				internalhelpers.VerifyUserDoesNotExist(t, ctx, sys, expectedUsername)
			}
		})
	}
}
