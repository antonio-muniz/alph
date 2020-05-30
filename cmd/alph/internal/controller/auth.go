package controller

import (
	"context"
	"time"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/cmd/alph/internal/database"
	"github.com/antonio-muniz/alph/cmd/alph/internal/database/memory"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/response"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

var (
	ErrIncorrectCredentials = errors.New("Incorrect credentials")
)

func Authenticate(ctx context.Context, components di.Container, request request.Authenticate) (response.Authenticate, error) {
	database := components.Get("database").(database.DB)
	subject, err := database.GetSubject(ctx, request.SubjectID)
	switch err {
	case nil:
	case memory.ErrSubjectNotFound:
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
			Subject:        request.SubjectID,
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
