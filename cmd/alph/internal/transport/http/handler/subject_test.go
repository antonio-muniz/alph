package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/database"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"
)

func TestCreateSubject(t *testing.T) {
	scenarios := []struct {
		description        string
		request            request.CreateSubject
		expectedStatusCode int
	}{
		{
			description: "creates_a_valid_subject",
			request: request.CreateSubject{
				SubjectID: "new.user@example.org",
				Password:  "hakunamatata",
			},
			expectedStatusCode: nethttp.StatusCreated,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			ctx := context.Background()
			components, err := internal.Components()
			require.NoError(t, err)
			router := http.Router(components)
			requestBody, err := json.Marshal(scenario.request)
			require.NoError(t, err)
			requestBodyReader := bytes.NewReader(requestBody)
			request, err := nethttp.NewRequest(nethttp.MethodPost, "/api/subjects", requestBodyReader)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			database := components.Get("database").(database.DB)
			subject, err := database.GetSubject(ctx, scenario.request.SubjectID)
			require.NoError(t, err)
			passwordMatch, err := password.Validate(scenario.request.Password, subject.HashedPassword)
			require.NoError(t, err)
			require.True(t, passwordMatch)
		})
	}
}
