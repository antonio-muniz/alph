package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/pkg/errors"
)

func AESEncrypt(message string, encryptionKey string) (string, error) {
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", errors.Wrap(err, "error creating AES cipher block")
	}
	messageBytes := []byte(message)
	encryptedMessageBytes := make([]byte, aes.BlockSize+len(messageBytes))
	initVector := encryptedMessageBytes[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, initVector)
	if err != nil {
		return "", errors.Wrap(err, "error generating initialization vector")
	}
	encrypter := cipher.NewCFBEncrypter(block, initVector)
	encrypter.XORKeyStream(encryptedMessageBytes[aes.BlockSize:], messageBytes)
	encryptedMessage := base64.RawURLEncoding.EncodeToString(encryptedMessageBytes)
	return encryptedMessage, nil
}

func AESDecrypt(encryptedMessage string, encryptionKey string) (string, error) {
	encryptedMessageBytes, err := base64.RawURLEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", errors.Wrap(err, "error decoding encrypted message")
	}
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", errors.Wrap(err, "error creating AES cipher block")
	}
	initVector := encryptedMessageBytes[:aes.BlockSize]
	encryptedMessageBytes = encryptedMessageBytes[aes.BlockSize:]
	decrypter := cipher.NewCFBDecrypter(block, initVector)
	decryptedMessageBytes := make([]byte, len(encryptedMessageBytes))
	decrypter.XORKeyStream(decryptedMessageBytes, encryptedMessageBytes)
	decryptedMessage := string(decryptedMessageBytes)
	return decryptedMessage, nil
}
