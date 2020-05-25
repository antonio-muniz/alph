package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/antonio-muniz/alph/pkg/models/token"
	"github.com/pkg/errors"
)

func Serialize(token token.Token) (string, error) {
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

func serializeHeader(header token.Header) (string, error) {
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize token header")
	}
	return encodeBase64URL(headerJSON), nil
}

func serializePayload(payload token.Payload) (string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize token payload")
	}
	return encodeBase64URL(payloadJSON), nil
}

func encodeBase64URL(bytes []byte) string {
	return base64.RawURLEncoding.EncodeToString(bytes)
}
