package controller

import (
	"context"
	"time"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/models/request"
	"github.com/antonio-muniz/alph/pkg/models/response"
	"github.com/antonio-muniz/alph/pkg/models/token"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

const (
	correctIdentity = "someone@example.org"
	correctPassword = "123456"
)

var (
	ErrIncorrectCredentials = errors.New("Incorrect credentials")
)

func Authenticate(ctx context.Context, components di.Container, request request.Authenticate) (response.Authenticate, error) {
	correctPasswordHash, err := password.Hash("123456")
	if err != nil {
		return response.Authenticate{}, errors.Wrap(err, "hashing correct password")
	}
	passwordCorrect, err := password.Validate(request.Password, correctPasswordHash)
	if err != nil {
		return response.Authenticate{}, errors.Wrap(err, "validating password")
	}
	if request.Identity != correctIdentity || !passwordCorrect {
		return response.Authenticate{}, ErrIncorrectCredentials
	}
	now := time.Now()
	token := token.Token{
		Header: token.Header{
			SignatureAlgorithm: "HS256",
			TokenType:          "JWT",
		},
		Payload: token.Payload{
			Audience:       "example.org",
			ExpirationTime: token.Timestamp(now.Add(30 * time.Minute)),
			IssuedAt:       token.Timestamp(now),
			Issuer:         "alph",
			Subject:        request.Identity,
		},
	}
	encodedToken, err := jwt.Serialize(token)
	if err != nil {
		return response.Authenticate{}, errors.Wrap(err, "serializing JWT")
	}
	config := components.Get("config").(config.Config)
	signedToken, err := jwt.Sign(encodedToken, "HS256", config.JwtSignatureKey)
	if err != nil {
		return response.Authenticate{}, errors.Wrap(err, "signing JWT")
	}
	accessToken, err := jwt.Encrypt(signedToken, config.JwtEncryptionPublicKey)
	if err != nil {
		return response.Authenticate{}, errors.Wrap(err, "encrypting JWT")
	}

	authResponse := response.Authenticate{AccessToken: accessToken}

	return authResponse, nil
}
