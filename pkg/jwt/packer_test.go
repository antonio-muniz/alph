package jwt_test

import (
	"testing"
	"time"

	"github.com/antonio-muniz/alph/pkg/jwt"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/encryption"
	"github.com/stretchr/testify/require"
)

func TestPackUnpack(t *testing.T) {
	scenarios := []struct {
		description   string
		token         jwt.Token
		signatureKey  string
		encryptionKey string
		decryptionKey string
	}{
		{
			description: "packs_and_unpacks_a_token",
			token: jwt.Token{
				Issuer:   "alph",
				Audience: "example.org",
				Subject:  "someone@example.org",
				ExpirationTime: jwt.Timestamp(
					time.Date(2020, time.May, 24, 20, 35, 37, 0, time.UTC),
				),
			},
			signatureKey:  "zLcwW6w2MEwS8RMzP71azVbQJyOK4fiV",
			encryptionKey: fixtures.PublicKey(),
			decryptionKey: fixtures.PrivateKey(),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			packSettings := jwt.PackSettings{
				SignatureKey:  scenario.signatureKey,
				EncryptionKey: scenario.encryptionKey,
			}
			packedToken, err := jwt.Pack(scenario.token, packSettings)
			require.NoError(t, err)
			unpackSettings := jwt.UnpackSettings{
				SignatureKey:  scenario.signatureKey,
				DecryptionKey: scenario.decryptionKey,
			}
			unpackedToken, err := jwt.Unpack(packedToken, unpackSettings)
			require.NoError(t, err)
			require.Equal(t, scenario.token, unpackedToken)
		})
	}
}
