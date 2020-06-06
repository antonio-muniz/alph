package controller

import (
	"context"
	"time"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/message"
	"github.com/antonio-muniz/alph/pkg/clock"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/pkg/errors"
)

var (
	ErrIncorrectCredentials = errors.New("Incorrect credentials")
)

func PasswordAuth(ctx context.Context, sys system.System, request message.PasswordAuthRequest) (message.PasswordAuthResponse, error) {
	user, err := findUser(ctx, sys, request.Username)
	switch err {
	case nil:
	case storage.ErrUserNotFound:
		return message.PasswordAuthResponse{}, ErrIncorrectCredentials
	default:
		return message.PasswordAuthResponse{}, errors.Wrap(err, "loading user")
	}
	passwordCorrect, err := password.Validate(request.Password, user.HashedPassword)
	if err != nil {
		return message.PasswordAuthResponse{}, errors.Wrap(err, "validating password")
	}
	if !passwordCorrect {
		return message.PasswordAuthResponse{}, ErrIncorrectCredentials
	}
	err = ensureClientExists(ctx, sys, request.ClientID, request.ClientSecret)
	switch err {
	case nil:
	case storage.ErrClientNotFound:
		return message.PasswordAuthResponse{}, ErrIncorrectCredentials
	default:
		return message.PasswordAuthResponse{}, errors.Wrap(err, "ensuring client exists")
	}

	accessToken, err := generateAccessToken(ctx, sys, request.Username)
	if err != nil {
		return message.PasswordAuthResponse{}, errors.Wrap(err, "generating access token")
	}

	authResponse := message.PasswordAuthResponse{AccessToken: accessToken}

	return authResponse, nil
}

func findUser(ctx context.Context, sys system.System, username string) (model.User, error) {
	database := sys.Get("database").(storage.Database)
	user, err := database.GetUser(ctx, username)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func ensureClientExists(
	ctx context.Context,
	sys system.System,
	clientID string,
	clientSecret string,
) error {
	if clientID != "the-client" {
		return storage.ErrClientNotFound
	}
	if clientSecret != "the-client-is-scared-of-the-dark" {
		return storage.ErrClientNotFound
	}
	return nil
}

func generateAccessToken(ctx context.Context, sys system.System, username string) (string, error) {
	clock := sys.Get("clock").(clock.Clock)
	now := clock.Now()
	token := jwt.Token{
		Audience:       "example.org",
		ExpirationTime: jwt.Timestamp(now.Add(30 * time.Minute)),
		Issuer:         "alph",
		Subject:        username,
	}
	config := sys.Get("config").(config.Config)
	accessToken, err := jwt.Pack(token, jwt.PackSettings{
		SignatureKey:  config.JWTSignatureKey,
		EncryptionKey: config.JWTEncryptionPublicKey,
	})
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
