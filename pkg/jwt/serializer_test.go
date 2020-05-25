package jwt_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/models/token"
)

func TestSerialize(t *testing.T) {
	t.Run("Serializes a token to string", func(t *testing.T) {
		token := token.Token{
			Header: token.TokenHeader{
				SignatureAlgorithm: "HS256",
				TokenType:          "JWT",
			},
			Payload: token.TokenPayload{
				Issuer:         "alph",
				Audience:       "example.org",
				Subject:        "someone@example.org",
				IssuedAt:       time.Date(2020, time.May, 24, 20, 00, 00, 000000000, time.UTC),
				ExpirationTime: time.Date(2020, time.May, 24, 20, 30, 00, 000000000, time.UTC),
			},
		}

		serializedToken, err := jwt.Serialize(token)
		require.NoError(t, err)
		require.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleGFtcGxlLm9yZyIsImV4cCI6IjIwMjAtMDUtMjRUMjA6MzA6MDBaIiwiaWF0IjoiMjAyMC0wNS0yNFQyMDowMDowMFoiLCJpc3MiOiJhbHBoIiwic3ViIjoic29tZW9uZUBleGFtcGxlLm9yZyJ9", serializedToken)
	})
}
