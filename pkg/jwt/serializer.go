package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/antonio-muniz/alph/pkg/models/token"
	"github.com/pkg/errors"
)

func Serialize(token token.Token) (string, error) {
	headerJSON, err := json.Marshal(map[string]interface{}{
		"alg": token.Header.SignatureAlgorithm,
		"typ": token.Header.TokenType,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize token header")
	}

	payloadJSON, err := json.Marshal(map[string]interface{}{
		"iss": token.Payload.Issuer,
		"aud": token.Payload.Audience,
		"sub": token.Payload.Subject,
		"iat": token.Payload.IssuedAt,
		"exp": token.Payload.ExpirationTime,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize token payload")
	}

	headerComponent := base64.RawURLEncoding.EncodeToString(headerJSON)
	payloadComponent := base64.RawURLEncoding.EncodeToString(payloadJSON)

	return strings.Join([]string{headerComponent, payloadComponent}, "."), nil
}
