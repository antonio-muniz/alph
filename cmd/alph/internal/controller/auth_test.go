package controller_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/cmd/alph/internal/controller"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/password"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/encryption"
)

func TestPasswordAuth(t *testing.T) {
	scenarios := []struct {
		description            string
		correctUsername        string
		correctPassword        string
		correctClientID        string
		correctClientSecret    string
		request                message.PasswordAuthRequest
		expectedError          error
		expectedIssuer         string
		expectedAudience       string
		expectedSubject        string
		expectedExpirationTime time.Time
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
			expectedError:          nil,
			expectedIssuer:         "alph",
			expectedAudience:       "example.org",
			expectedSubject:        "someone@example.org",
			expectedExpirationTime: time.Now().Add(30 * time.Minute),
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
			if scenario.expectedError == nil {
				config := sys.Get("config").(config.Config)
				unpackedToken, err := jwt.Unpack(
					response.AccessToken,
					jwt.UnpackSettings{
						DecryptionKey: fixtures.PrivateKey(),
						SignatureKey:  config.JWTSignatureKey,
					},
				)
				require.NoError(t, err)
				require.Equal(t, scenario.expectedIssuer, unpackedToken.Issuer)
				require.Equal(t, scenario.expectedAudience, unpackedToken.Audience)
				require.Equal(t, scenario.expectedSubject, unpackedToken.Subject)
				require.WithinDuration(t,
					scenario.expectedExpirationTime,
					time.Time(unpackedToken.ExpirationTime),
					1*time.Second,
				)
			}
		})
	}
}
