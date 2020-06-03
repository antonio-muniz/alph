package controller_test

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/test/internalhelpers"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
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
			internalhelpers.VerifyUserExists(t,
				ctx,
				sys,
				scenario.request.Username,
				scenario.request.Password,
			)
		})
	}
}
