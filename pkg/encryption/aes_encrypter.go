package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/pkg/errors"
)

var ErrBadAESKeyLength = errors.New("AES key is not 32 bytes long")
var ErrBadEncryptedMessageEncoding = errors.New("encrypted message is not in base64")

func AESEncrypt(message string, encryptionKey string) (string, error) {
	encryptionKeyBytes := []byte(encryptionKey)
	err := validateAESKey(encryptionKeyBytes)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(encryptionKeyBytes)
	if err != nil {
		return "", errors.Wrap(err, "creating AES cipher block")
	}
	messageBytes := []byte(message)
	encryptedMessageBytes := make([]byte, aes.BlockSize+len(messageBytes))
	initVector := encryptedMessageBytes[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, initVector)
	if err != nil {
		return "", errors.Wrap(err, "generating initialization vector")
	}
	encrypter := cipher.NewCFBEncrypter(block, initVector)
	encrypter.XORKeyStream(encryptedMessageBytes[aes.BlockSize:], messageBytes)
	encryptedMessage := base64.RawURLEncoding.EncodeToString(encryptedMessageBytes)
	return encryptedMessage, nil
}

func AESDecrypt(encryptedMessage string, decryptionKey string) (string, error) {
	decryptionKeyBytes := []byte(decryptionKey)
	err := validateAESKey(decryptionKeyBytes)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(decryptionKeyBytes)
	if err != nil {
		return "", errors.Wrap(err, "creating AES cipher block")
	}
	encryptedMessageBytes, err := base64.RawURLEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", ErrBadEncryptedMessageEncoding
	}
	initVector := encryptedMessageBytes[:aes.BlockSize]
	encryptedMessageBytes = encryptedMessageBytes[aes.BlockSize:]
	decrypter := cipher.NewCFBDecrypter(block, initVector)
	decryptedMessageBytes := make([]byte, len(encryptedMessageBytes))
	decrypter.XORKeyStream(decryptedMessageBytes, encryptedMessageBytes)
	decryptedMessage := string(decryptedMessageBytes)
	return decryptedMessage, nil
}

func validateAESKey(keyBytes []byte) error {
	if len(keyBytes) != 32 {
		return ErrBadAESKeyLength
	}
	return nil
}
