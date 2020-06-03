package controller

import (
	"context"
	"time"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
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
	database := sys.Get("database").(storage.Database)
	user, err := database.GetUser(ctx, request.Username)
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
	if request.ClientID != "the-client" {
		return message.PasswordAuthResponse{}, ErrIncorrectCredentials
	}
	if request.ClientSecret != "the-client-is-scared-of-the-dark" {
		return message.PasswordAuthResponse{}, ErrIncorrectCredentials
	}
	clock := sys.Get("clock").(clock.Clock)
	now := clock.Now()
	token := jwt.Token{
		Audience:       "example.org",
		ExpirationTime: jwt.Timestamp(now.Add(30 * time.Minute)),
		Issuer:         "alph",
		Subject:        request.Username,
	}
	config := sys.Get("config").(config.Config)
	accessToken, err := jwt.Pack(token, jwt.PackSettings{
		SignatureKey:  config.JWTSignatureKey,
		EncryptionKey: config.JWTEncryptionPublicKey,
	})

	authResponse := message.PasswordAuthResponse{AccessToken: accessToken}

	return authResponse, nil
}
