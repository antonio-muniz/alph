package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
)

func Sign(encodedToken string, signingKey string) (string, error) {
	signature, err := generateSignature(encodedToken, signingKey)
	if err != nil {
		return "", err
	}
	signedToken := fmt.Sprintf("%s.%s", encodedToken, signature)
	return signedToken, nil
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
