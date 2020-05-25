package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/pkg/errors"
)

func Encrypt(message string, encryptionKey string) (string, error) {
	publicKey, err := parsePublicKey(encryptionKey)
	if err != nil {
		return "", err
	}
	encryptedMessage, err := encryptMessage(message, publicKey)
	if err != nil {
		return "", err
	}
	return encryptedMessage, nil
}

func Decrypt(encryptedMessage string, decryptionKey string) (string, error) {
	privateKey, err := parsePrivateKey(decryptionKey)
	if err != nil {
		return "", err
	}
	decryptedMessage, err := decryptMessage(encryptedMessage, privateKey)
	if err != nil {
		return "", err
	}
	return decryptedMessage, nil
}

func parsePublicKey(key string) (*rsa.PublicKey, error) {
	pemBlock, _ := pem.Decode([]byte(key))
	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse encryption key")
	}
	rsaPublicKey, isRSAPublicKey := publicKey.(*rsa.PublicKey)
	if !isRSAPublicKey {
		return nil, errors.New("encryption key is not a RSA public key")
	}
	return rsaPublicKey, nil
}

func encryptMessage(message string, publicKey *rsa.PublicKey) (string, error) {
	encryptedMessageBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(message),
		nil,
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to encrypt message")
	}
	encryptedMessage := base64.RawURLEncoding.EncodeToString(encryptedMessageBytes)
	return encryptedMessage, nil
}

func parsePrivateKey(key string) (*rsa.PrivateKey, error) {
	pemBlock, _ := pem.Decode([]byte(key))
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse decryption key")
	}
	rsaPrivateKey, isRSAPrivateKey := privateKey.(*rsa.PrivateKey)
	if !isRSAPrivateKey {
		return nil, errors.New("decryption key is not a RSA private key")
	}
	return rsaPrivateKey, nil
}

func decryptMessage(encryptedMessage string, privateKey *rsa.PrivateKey) (string, error) {
	encryptedMessageBytes, err := base64.RawURLEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", errors.Wrap(err, "encrypted message is not base64 encoded")
	}
	decryptedMessageBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		encryptedMessageBytes,
		nil,
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt message")
	}
	decryptedMessage := string(decryptedMessageBytes)
	return decryptedMessage, nil
}
