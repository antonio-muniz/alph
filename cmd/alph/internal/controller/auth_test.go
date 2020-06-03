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
	"github.com/antonio-muniz/alph/pkg/system"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/encryption"
)

func TestPasswordAuth(t *testing.T) {
	scenarios := []struct {
		description         string
		correctUsername     string
		correctPassword     string
		correctClientID     string
		correctClientSecret string
		request             message.PasswordAuthRequest
		expectedError       error
		expectedToken       jwt.Token
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
			expectedToken: jwt.Token{
				Issuer:         "alph",
				Audience:       "example.org",
				Subject:        "someone@example.org",
				ExpirationTime: jwt.Timestamp(time.Now().Add(30 * time.Minute)),
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
			sys := initializeSystem(t, ctx)
			user := model.User{
				Username:       scenario.correctUsername,
				HashedPassword: hashPassword(t, scenario.correctPassword),
			}
			createUser(t, ctx, sys, user)
			response, err := controller.PasswordAuth(ctx, sys, scenario.request)
			require.Equal(t, scenario.expectedError, err)
			if scenario.expectedError == nil {
				verifyAccessTokenWithLooseExpiration(t, sys, scenario.expectedToken, response.AccessToken)
			}
		})
	}
}

func initializeSystem(t *testing.T, ctx context.Context) system.System {
	sys, err := internal.System(ctx)
	require.NoError(t, err)
	return sys
}

func hashPassword(t *testing.T, plainPassword string) string {
	hashedPassword, err := password.Hash(plainPassword)
	require.NoError(t, err)
	return hashedPassword
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

func verifyAccessTokenWithLooseExpiration(
	t *testing.T,
	sys system.System,
	expectedToken jwt.Token,
	accessToken string,
) {
	config := sys.Get("config").(config.Config)
	unpackedToken, err := jwt.Unpack(
		accessToken,
		jwt.UnpackSettings{
			DecryptionKey: fixtures.PrivateKey(),
			SignatureKey:  config.JWTSignatureKey,
		},
	)
	require.NoError(t, err)
	require.Equal(t, expectedToken.Issuer, unpackedToken.Issuer)
	require.Equal(t, expectedToken.Audience, unpackedToken.Audience)
	require.Equal(t, expectedToken.Subject, unpackedToken.Subject)
	require.WithinDuration(t,
		time.Time(expectedToken.ExpirationTime),
		time.Time(unpackedToken.ExpirationTime),
		1*time.Second,
	)
}
