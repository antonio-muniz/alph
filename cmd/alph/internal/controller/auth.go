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

func Authenticate(ctx context.Context, sys system.System, request request.Authenticate) (response.Authenticate, error) {
	database := sys.Get("database").(storage.Database)
	subject, err := database.GetSubject(ctx, request.Username)
	switch err {
	case nil:
	case storage.ErrSubjectNotFound:
		return response.Authenticate{}, ErrIncorrectCredentials
	default:
		return response.Authenticate{}, errors.Wrap(err, "loading subject")
	}
	passwordCorrect, err := password.Validate(request.Password, subject.HashedPassword)
	if err != nil {
		return response.Authenticate{}, errors.Wrap(err, "validating password")
	}
	if !passwordCorrect {
		return response.Authenticate{}, ErrIncorrectCredentials
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

	authResponse := response.Authenticate{AccessToken: accessToken}

	return authResponse, nil
}
