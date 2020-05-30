package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var ErrIncorrectSignature = errors.New("incorrect signature")

func Sign(encodedToken string, signingKey string) (string, error) {
	signature, err := generateSignature(encodedToken, signingKey)
	if err != nil {
		return "", err
	}
	signedToken := fmt.Sprintf("%s.%s", encodedToken, signature)
	return signedToken, nil
}

func CheckSignature(signedToken string, signingKey string) error {
	separationIndex := strings.LastIndex(signedToken, ".")
	unsignedToken := signedToken[:separationIndex]
	expectedSignedToken, err := Sign(unsignedToken, signingKey)
	if err != nil {
		return err
	}
	if expectedSignedToken != signedToken {
		return ErrIncorrectSignature
	}
	return nil
}

func generateSignature(encodedToken string, signingKey string) (string, error) {
	hashFunction := hmac.New(sha256.New, []byte(signingKey))
	_, err := hashFunction.Write([]byte(encodedToken))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate HMAC-SHA256 signature")
	}
	signature := hashFunction.Sum(nil)
	encodedSignature := base64.RawURLEncoding.EncodeToString(signature)
	return encodedSignature, nil
}
