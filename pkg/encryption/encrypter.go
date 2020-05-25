package encryption

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

func Decrypt(encryptedMessage string, encryptionAlgorithm string, decryptionKey string) (string, error) {
	if encryptionAlgorithm != "RSA" {
		return "", fmt.Errorf("unsupported encryption algorithm '%s'", encryptionAlgorithm)
	}
	pemBlock, _ := pem.Decode([]byte(decryptionKey))
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse decryption key")
	}
	rsaPrivateKey, isRSAPrivateKey := privateKey.(*rsa.PrivateKey)
	if !isRSAPrivateKey {
		return "", errors.New("decryption key is not a RSA private key")
	}
	encryptedMessageBytes, err := base64.RawURLEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", errors.Wrap(err, "encrypted message is not base64 encoded")
	}
	decryptedMessageBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPrivateKey,
		encryptedMessageBytes,
		nil,
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt message")
	}
	decryptedMessage := string(decryptedMessageBytes)
	return decryptedMessage, nil
}
