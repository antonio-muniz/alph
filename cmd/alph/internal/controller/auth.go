package controller

import (
	"context"
	"time"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/response"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/pkg/errors"
)

var (
	ErrIncorrectCredentials = errors.New("Incorrect credentials")
)

func PasswordAuth(ctx context.Context, sys system.System, request request.PasswordAuth) (response.PasswordAuth, error) {
	database := sys.Get("database").(storage.Database)
	user, err := database.GetUser(ctx, request.Username)
	switch err {
	case nil:
	case storage.ErrUserNotFound:
		return response.PasswordAuth{}, ErrIncorrectCredentials
	default:
		return response.PasswordAuth{}, errors.Wrap(err, "loading user")
	}
	passwordCorrect, err := password.Validate(request.Password, user.HashedPassword)
	if err != nil {
		return response.PasswordAuth{}, errors.Wrap(err, "validating password")
	}
	if !passwordCorrect {
		return response.PasswordAuth{}, ErrIncorrectCredentials
	}
	now := time.Now()
	token := jwt.Token{
		Header: jwt.Header{
			SignatureAlgorithm: "HS256",
			TokenType:          "JWT",
		},
		Payload: jwt.Payload{
			Audience:       "example.org",
			ExpirationTime: jwt.Timestamp(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.Timestamp(now),
			Issuer:         "alph",
			Subject:        request.Username,
		},
	}
	config := sys.Get("config").(config.Config)
	accessToken, err := jwt.Pack(token, jwt.PackSettings{
		SignatureKey:  config.JWTSignatureKey,
		EncryptionKey: config.JWTEncryptionPublicKey,
	})

	authResponse := response.PasswordAuth{AccessToken: accessToken}

	return authResponse, nil
}
