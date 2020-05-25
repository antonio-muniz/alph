package jwt_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"

	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/test/fixtures/encryption"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	scenarios := []struct {
		description           string
		message               string
		encryptionAlgorithm   string
		encryptionKey         string
		expectedDecryptionKey string
	}{
		{
			description:           "encrypts_a_message_using_RSA_algorithm",
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

			pemBlock, _ := pem.Decode([]byte(scenario.expectedDecryptionKey))
			privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
			require.NoError(t, err)
			rsaPrivateKey, isRSAPrivateKey := privateKey.(*rsa.PrivateKey)
			if !isRSAPrivateKey {
				t.Fatal("expected decryption key is not a RSA private key")
			}
			encryptedMessageBytes, err := base64.RawURLEncoding.DecodeString(encryptedMessage)
			require.NoError(t, err)
			decryptedMessageBytes, err := rsa.DecryptOAEP(
				sha256.New(),
				rand.Reader,
				rsaPrivateKey,
				encryptedMessageBytes,
				nil,
			)
			require.NoError(t, err)
			decryptedMessage := string(decryptedMessageBytes)
			require.Equal(t, scenario.message, decryptedMessage)
		})
	}
}
