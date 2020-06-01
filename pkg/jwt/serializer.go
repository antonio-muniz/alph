package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Serialize(token Token) (string, error) {
	serializedHeader, err := serializeHeader()
	if err != nil {
		return "", err
	}
	serializedToken, err := serializeToken(token)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", serializedHeader, serializedToken), nil
}

func Deserialize(serializedToken string) (Token, error) {
	tokenParts := strings.SplitN(serializedToken, ".", 2)
	serializedHeader := tokenParts[0]
	_, err := deserializeHeader(serializedHeader)
	if err != nil {
		return Token{}, err
	}
	serializedPayload := tokenParts[1]
	token, err := deserializeToken(serializedPayload)
	if err != nil {
		return Token{}, err
	}
	return token, nil
}

func serializeHeader() (string, error) {
	headerJSON, err := json.Marshal(Header{
		SignatureAlgorithm: "HS256",
		TokenType:          "JWT",
	})
	if err != nil {
		return "", errors.Wrap(err, "serializing token header")
	}
	encodedHeader := base64.RawURLEncoding.EncodeToString(headerJSON)
	return encodedHeader, nil
}

func deserializeHeader(serializedHeader string) (Header, error) {
	headerJSON, err := base64.RawURLEncoding.DecodeString(serializedHeader)
	if err != nil {
		return Header{}, errors.Wrap(err, "decoding token header")
	}
	var header Header
	err = json.Unmarshal(headerJSON, &header)
	if err != nil {
		return Header{}, errors.Wrap(err, "deserializing token header")
	}
	return header, nil
}

func serializeToken(payload Token) (string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "serializing token payload")
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadJSON)
	return encodedPayload, nil
}

func deserializeToken(serializedToken string) (Token, error) {
	payloadJSON, err := base64.RawURLEncoding.DecodeString(serializedToken)
	if err != nil {
		return Token{}, errors.Wrap(err, "decoding token payload")
	}
	var payload Token
	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		return Token{}, errors.Wrap(err, "deserializing token payload")
	}
	return payload, nil
}
