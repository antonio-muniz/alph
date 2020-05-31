package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	scenarios := []struct {
		description        string
		correctSubjectID   string
		correctPassword    string
		request            request.Authenticate
		expectedStatusCode int
	}{
		{
			description:      "authenticates_subject_with_correct_password",
			correctSubjectID: "someone@example.org",
			correctPassword:  "123456",
			request: request.Authenticate{
				SubjectID: "someone@example.org",
				Password:  "123456",
			},
			expectedStatusCode: nethttp.StatusOK,
		},
		{
			description:      "does_not_authenticate_subject_with_incorrect_password",
			correctSubjectID: "someone@example.org",
			correctPassword:  "123456",
			request: request.Authenticate{
				SubjectID: "someone@example.org",
				Password:  "654321",
			},
			expectedStatusCode: nethttp.StatusForbidden,
		},
		{
			description:      "does_not_authenticate_unknown_subject",
			correctSubjectID: "someone@example.org",
			correctPassword:  "123456",
			request: request.Authenticate{
				SubjectID: "someone@example.org",
				Password:  "654321",
			},
			expectedStatusCode: nethttp.StatusForbidden,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			ctx := context.Background()
			sys, err := internal.System()
			require.NoError(t, err)
			hashedCorrectPassword, err := password.Hash(scenario.correctPassword)
			require.NoError(t, err)
			database := sys.Get("database").(storage.Database)
			subject := auth.Subject{
				ID:             scenario.correctSubjectID,
				HashedPassword: hashedCorrectPassword,
			}
			err = database.CreateSubject(ctx, subject)
			require.NoError(t, err)
			router := http.Router(sys)
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
