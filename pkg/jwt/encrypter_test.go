package jwt_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/jwt"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/encryption"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecrypt(t *testing.T) {
	scenarios := []struct {
		description           string
		signedToken           string
		encryptionKey         string
		expectedDecryptionKey string
	}{
		{
			description:           "encrypt_then_decrypt_a_signed_token",
			signedToken:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhbHBoIiwic3ViIjoic29tZW9uZUBleGFtcGxlLm9yZyIsImF1ZCI6ImV4YW1wbGUub3JnIiwiaWF0IjoiMjAyMC0wNS0yNFQyMDowMDowMFoiLCJleHAiOiIyMDIwLTA1LTI0VDIwOjMwOjAwWiJ9.AVC8mWAWEkQYYeduwnQVGyaOUXHKpQkbx4GT-iv7bOY",
			encryptionKey:         fixtures.PublicKey(),
			expectedDecryptionKey: fixtures.PrivateKey(),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			encryptedToken, err := jwt.Encrypt(scenario.signedToken, scenario.encryptionKey)
			require.NoError(t, err)
			decryptedToken, err := jwt.Decrypt(encryptedToken, scenario.expectedDecryptionKey)
			require.NoError(t, err)
			require.Equal(t, scenario.signedToken, decryptedToken)
		})
	}
}
