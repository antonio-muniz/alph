package controller_test

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/test/internalhelpers"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/antonio-muniz/alph/pkg/system"
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
			sys := internalhelpers.InitializeSystem(t, ctx)
			err := controller.NewUser(ctx, sys, scenario.request)
			require.Equal(t, scenario.expectedError, err)
			verifyUser(t, ctx, sys, scenario.request.Username, scenario.request.Password)
		})
	}
}

func verifyUser(
	t *testing.T,
	ctx context.Context,
	sys system.System,
	expectedUsername string,
	expectedPassword string,
) {
	database := sys.Get("database").(storage.Database)
	user, err := database.GetUser(ctx, expectedUsername)
	require.NoError(t, err)
	passwordMatch, err := password.Validate(expectedPassword, user.HashedPassword)
	require.NoError(t, err)
	require.True(t, passwordMatch)
}
