package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	scenarios := []struct {
		description        string
		request            request.NewUser
		expectedStatusCode int
	}{
		{
			description: "creates_a_valid_user",
			request: request.NewUser{
				Username: "new.user@example.org",
				Password: "hakunamatata",
			},
			expectedStatusCode: nethttp.StatusCreated,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			ctx := context.Background()
			sys, err := internal.System()
			require.NoError(t, err)
			router := http.Router(sys)
			requestBody, err := json.Marshal(scenario.request)
			require.NoError(t, err)
			requestBodyReader := bytes.NewReader(requestBody)
			request, err := nethttp.NewRequest(nethttp.MethodPost, "/api/users", requestBodyReader)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			database := sys.Get("database").(storage.Database)
			user, err := database.GetUser(ctx, scenario.request.Username)
			require.NoError(t, err)
			passwordMatch, err := password.Validate(scenario.request.Password, user.HashedPassword)
			require.NoError(t, err)
			require.True(t, passwordMatch)
		})
	}
}
