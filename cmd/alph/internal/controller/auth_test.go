package controller_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/test/internalhelpers"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/antonio-muniz/alph/test/helpers"
)

func TestPasswordAuth(t *testing.T) {
	scenarios := []struct {
		description           string
		correctUsername       string
		correctPassword       string
		correctClientID       string
		correctClientSecret   string
		request               message.PasswordAuthRequest
		expectedError         error
		expectedUnpackedToken jwt.Token
	}{
		{
			description:         "authenticates_user_with_correct_password_and_existing_client",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			request: message.PasswordAuthRequest{
				Username:     "someone@example.org",
				Password:     "123456",
				ClientID:     "the-client",
				ClientSecret: "the-client-is-scared-of-the-dark",
			},
			expectedError: nil,
			expectedUnpackedToken: jwt.Token{
				Issuer:         "alph",
				Audience:       "example.org",
				Subject:        "someone@example.org",
				ExpirationTime: jwt.Timestamp(helpers.Now().Add(30 * time.Minute)),
			},
		},
		{
			description:         "does_not_authenticate_unknown_user",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			request: message.PasswordAuthRequest{
				Username:     "someone.else@example.org",
				Password:     "123456",
				ClientID:     "the-client",
				ClientSecret: "the-client-is-scared-of-the-dark",
			},
			expectedError: controller.ErrIncorrectCredentials,
		},
		{
			description:         "does_not_authenticate_user_with_incorrect_password",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			request: message.PasswordAuthRequest{
				Username:     "someone@example.org",
				Password:     "654321",
				ClientID:     "the-client",
				ClientSecret: "the-client-is-scared-of-the-dark",
			},
			expectedError: controller.ErrIncorrectCredentials,
		},
		{
			description:         "does_not_authenticate_user_on_unknown_client",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			request: message.PasswordAuthRequest{
				Username:     "someone@example.org",
				Password:     "123456",
				ClientID:     "the-client-no-one-has-ever-heard-of",
				ClientSecret: "the-client-is-scared-of-the-dark",
			},
			expectedError: controller.ErrIncorrectCredentials,
		},
		{
			description:         "does_not_authenticate_user_with_incorrect_client_secret",
			correctUsername:     "someone@example.org",
			correctPassword:     "123456",
			correctClientID:     "the-client",
			correctClientSecret: "the-client-is-scared-of-the-dark",
			request: message.PasswordAuthRequest{
				Username:     "someone@example.org",
				Password:     "123456",
				ClientID:     "the-client",
				ClientSecret: "the-client-is-scared-of-butterflies",
			},
			expectedError: controller.ErrIncorrectCredentials,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			ctx := context.Background()
			sys := internalhelpers.InitializeSystem(t, ctx)
			user := model.User{
				Username:       scenario.correctUsername,
				HashedPassword: helpers.HashPassword(t, scenario.correctPassword),
			}
			createUser(t, ctx, sys, user)
			response, err := controller.PasswordAuth(ctx, sys, scenario.request)
			require.Equal(t, scenario.expectedError, err)
			if scenario.expectedError == nil {
				internalhelpers.VerifyAccessToken(t,
					sys,
					scenario.expectedUnpackedToken,
					response.AccessToken,
				)
			}
		})
	}
}

func createUser(
	t *testing.T,
	ctx context.Context,
	sys system.System,
	user model.User,
) {
	database := sys.Get("database").(storage.Database)
	err := database.CreateUser(ctx, user)
	require.NoError(t, err)
}
