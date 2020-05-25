package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/pkg/errors"
)

func Encrypt(message string, encryptionAlgorithm string, encryptionKey string) (string, error) {
	if encryptionAlgorithm != "RSA" {
		return "", fmt.Errorf("unsupported encryption algorithm '%s'", encryptionAlgorithm)
	}
	pemBlock, _ := pem.Decode([]byte(encryptionKey))
	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse encryption key")
	}
	rsaPublicKey, isRSAPublicKey := publicKey.(*rsa.PublicKey)
	if !isRSAPublicKey {
		return "", errors.New("encryption key is not a RSA public key")
	}
	encryptedMessageBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		[]byte(message),
		nil,
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to encrypt message")
	}
	encryptedMessage := base64.RawURLEncoding.EncodeToString(encryptedMessageBytes)
	return encryptedMessage, nil
}
