package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func Serialize(token Token) (string, error) {
	serializedHeader, err := serializeHeader(token.Header)
	if err != nil {
		return "", err
	}
	serializedPayload, err := serializePayload(token.Payload)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", serializedHeader, serializedPayload), nil
}

func serializeHeader(header Header) (string, error) {
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize token header")
	}
	encodedHeader := base64.RawURLEncoding.EncodeToString(headerJSON)
	return encodedHeader, nil
}

func serializePayload(payload Payload) (string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize token payload")
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadJSON)
	return encodedPayload, nil
}
