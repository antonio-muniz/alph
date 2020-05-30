package handler_test

import (
	"bytes"
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	scenarios := []struct {
		description        string
		request            request.Authenticate
		expectedStatusCode int
	}{
		{
			description: "authenticates_subject_with_correct_password",
			request: request.Authenticate{
				SubjectID: "someone@example.org",
				Password:  "123456",
			},
			expectedStatusCode: nethttp.StatusOK,
		},
		{
			description: "does_not_authenticate_subject_with_incorrect_password",
			request: request.Authenticate{
				SubjectID: "someone@example.org",
				Password:  "654321",
			},
			expectedStatusCode: nethttp.StatusForbidden,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			components, err := internal.Components()
			require.NoError(t, err)
			router := http.Router(components)
			requestBody, err := json.Marshal(scenario.request)
			require.NoError(t, err)
			requestBodyReader := bytes.NewReader(requestBody)
			request, err := nethttp.NewRequest(nethttp.MethodPost, "/api/auth", requestBodyReader)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
		})
	}
}
