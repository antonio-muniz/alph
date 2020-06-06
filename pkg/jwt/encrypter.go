package jwt

import (
	"crypto/rand"
	"io"
	"strings"

	"github.com/antonio-muniz/alph/pkg/encryption"
	"github.com/pkg/errors"
)

func Encrypt(signedToken string, encryptionKey string) (string, error) {
	aesEncryptionKeyBytes := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, aesEncryptionKeyBytes)
	if err != nil {
		return "", errors.WithStack(err)
	}
	aesEncryptionKey := string(aesEncryptionKeyBytes)
	aesEncryptedToken, err := encryption.AESEncrypt(signedToken, aesEncryptionKey)
	if err != nil {
		return "", err
	}
	encryptedAESKey, err := encryption.RSAEncrypt(aesEncryptionKey, encryptionKey)
	if err != nil {
		return "", err
	}
	encryptedToken := strings.Join([]string{aesEncryptedToken, encryptedAESKey}, ".")
	return encryptedToken, nil
}

func Decrypt(encryptedToken string, decryptionKey string) (string, error) {
	encryptedTokenParts := strings.SplitN(encryptedToken, ".", 2)
	aesEncryptedToken := encryptedTokenParts[0]
	encryptedAESKey := encryptedTokenParts[1]
	aesEncryptionKey, err := encryption.RSADecrypt(encryptedAESKey, decryptionKey)
	if err != nil {
		return "", err
	}
	signedToken, err := encryption.AESDecrypt(aesEncryptedToken, aesEncryptionKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
