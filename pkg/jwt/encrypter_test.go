package jwt_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/test/fixtures/encryption"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecrypt(t *testing.T) {
	scenarios := []struct {
		description           string
		message               string
		encryptionAlgorithm   string
		encryptionKey         string
		expectedDecryptionKey string
	}{
		{
			description:           "encrypts_then_decrypts_a_message_using_RSA_algorithm",
			message:               "shhh-this-is-secret",
			encryptionAlgorithm:   "RSA",
			encryptionKey:         encryption.PublicKey(),
			expectedDecryptionKey: encryption.PrivateKey(),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			encryptedMessage, err := jwt.Encrypt(
				scenario.message,
				scenario.encryptionAlgorithm,
				scenario.encryptionKey,
			)
			require.NoError(t, err)

			decryptedMessage, err := jwt.Decrypt(
				encryptedMessage,
				scenario.encryptionAlgorithm,
				scenario.expectedDecryptionKey,
			)
			require.NoError(t, err)
			require.Equal(t, scenario.message, decryptedMessage)
		})
	}
}
