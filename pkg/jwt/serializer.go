package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Serialize(token OldToken) (string, error) {
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

func Deserialize(serializedToken string) (OldToken, error) {
	tokenParts := strings.SplitN(serializedToken, ".", 2)
	serializedHeader := tokenParts[0]
	header, err := deserializeHeader(serializedHeader)
	if err != nil {
		return OldToken{}, err
	}
	serializedPayload := tokenParts[1]
	payload, err := deserializePayload(serializedPayload)
	if err != nil {
		return OldToken{}, err
	}
	token := OldToken{
		Header:  header,
		Payload: payload,
	}
	return token, nil
}

func serializeHeader(header Header) (string, error) {
	headerJSON, err := json.Marshal(header)
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

func serializePayload(payload Token) (string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "serializing token payload")
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadJSON)
	return encodedPayload, nil
}

func deserializePayload(serializedPayload string) (Token, error) {
	fmt.Println(serializedPayload)
	payloadJSON, err := base64.RawURLEncoding.DecodeString(serializedPayload)
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
