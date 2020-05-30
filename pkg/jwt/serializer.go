package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

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

func Deserialize(serializedToken string) (Token, error) {
	tokenParts := strings.SplitN(serializedToken, ".", 2)
	serializedHeader := tokenParts[0]
	header, err := deserializeHeader(serializedHeader)
	if err != nil {
		return Token{}, err
	}
	serializedPayload := tokenParts[1]
	payload, err := deserializePayload(serializedPayload)
	if err != nil {
		return Token{}, err
	}
	token := Token{
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

func serializePayload(payload Payload) (string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "serializing token payload")
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadJSON)
	return encodedPayload, nil
}

func deserializePayload(serializedPayload string) (Payload, error) {
	fmt.Println(serializedPayload)
	payloadJSON, err := base64.RawURLEncoding.DecodeString(serializedPayload)
	if err != nil {
		return Payload{}, errors.Wrap(err, "decoding token payload")
	}
	var payload Payload
	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		return Payload{}, errors.Wrap(err, "deserializing token payload")
	}
	return payload, nil
}
