package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"

	"github.com/pkg/errors"
)

func Sign(encodedToken string, signingAlgorithm string, signingKey string) (string, error) {
	signature, err := generateSignature(encodedToken, signingAlgorithm, signingKey)
	if err != nil {
		return "", err
	}
	signedToken := fmt.Sprintf("%s.%s", encodedToken, signature)
	return signedToken, nil
}

func generateSignature(encodedToken string, signingAlgorithm string, signingKey string) (string, error) {
	hashFunction, err := createHashFunction(signingAlgorithm, signingKey)
	if err != nil {
		return "", err
	}
	_, err = hashFunction.Write([]byte(encodedToken))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate HMAC-SHA256 signature")
	}
	signature := hashFunction.Sum(nil)
	encodedSignature := base64.RawURLEncoding.EncodeToString(signature)
	return encodedSignature, nil
}

func createHashFunction(signingAlgorithm string, signingKey string) (hash.Hash, error) {
	switch signingAlgorithm {
	case "HS256":
		return hmac.New(sha256.New, []byte(signingKey)), nil
	default:
		return nil, fmt.Errorf("unsupported signing algorithm '%s'", signingAlgorithm)
	}
}
