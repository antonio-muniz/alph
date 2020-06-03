package controller_test

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	scenarios := []struct {
		description   string
		request       message.NewUserRequest
		expectedError error
	}{
		{
			description: "creates_a_user",
			request: message.NewUserRequest{
				Username: "new.user@example.org",
				Password: "hakunamatata",
			},
			expectedError: nil,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			ctx := context.Background()
			sys, err := internal.System(ctx)
			require.NoError(t, err)
			err = controller.NewUser(ctx, sys, scenario.request)
			require.Equal(t, scenario.expectedError, err)
			database := sys.Get("database").(storage.Database)
			user, err := database.GetUser(ctx, scenario.request.Username)
			require.NoError(t, err)
			passwordMatch, err := password.Validate(scenario.request.Password, user.HashedPassword)
			require.NoError(t, err)
			require.True(t, passwordMatch)
		})
	}
}
