package controller_test

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
)

func TestPasswordAuth(t *testing.T) {
	scenarios := []struct {
		description         string
		correctUsername     string
		correctPassword     string
		correctClientID     string
		correctClientSecret string
		request             message.PasswordAuthRequest
		expectedToken       bool
		expectedError       error
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
			expectedToken: true,
			expectedError: nil,
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
			expectedToken: false,
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
			expectedToken: false,
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
			expectedToken: false,
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
			expectedToken: false,
			expectedError: controller.ErrIncorrectCredentials,
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
			user := model.User{
				Username:       scenario.correctUsername,
				HashedPassword: hashedCorrectPassword,
			}
			err = database.CreateUser(ctx, user)
			require.NoError(t, err)
			response, err := controller.PasswordAuth(ctx, sys, scenario.request)
			require.Equal(t, scenario.expectedError, err)
			require.Equal(t, scenario.expectedToken, len(response.AccessToken) > 0)
		})
	}
}
