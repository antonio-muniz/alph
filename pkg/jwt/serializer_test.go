package jwt_test

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/alph/pkg/jwt"
)

func TestSerialize(t *testing.T) {
	scenarios := []struct {
		description     string
		token           jwt.Token
		expectedHeader  map[string]interface{}
		expectedPayload map[string]interface{}
	}{
		{
			description: "serializes_a_token_to_string",
			token: jwt.Token{
				Header: jwt.Header{
					SignatureAlgorithm: "HS256",
					TokenType:          "JWT",
				},
				Payload: jwt.Payload{
					Issuer:   "alph",
					Audience: "example.org",
					Subject:  "someone@example.org",
					IssuedAt: jwt.Timestamp(
						time.Date(2020, time.May, 24, 20, 05, 37, 165098132, time.UTC),
					),
					ExpirationTime: jwt.Timestamp(
						time.Date(2020, time.May, 24, 20, 35, 37, 165098132, time.UTC),
					),
				},
			},
			expectedHeader: map[string]interface{}{
				"alg": "HS256",
				"typ": "JWT",
			},
			expectedPayload: map[string]interface{}{
				"iss": "alph",
				"aud": "example.org",
				"sub": "someone@example.org",
				"iat": float64(1590350737),
				"exp": float64(1590352537),
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			serializedToken, err := jwt.Serialize(scenario.token)
			require.NoError(t, err)

			tokenComponents := strings.SplitN(serializedToken, ".", 2)
			require.Len(t, tokenComponents, 2)

			header := deserializeTokenComponent(t, tokenComponents[0])
			require.Equal(t, scenario.expectedHeader, header)

			payload := deserializeTokenComponent(t, tokenComponents[1])
			require.Equal(t, scenario.expectedPayload, payload)
		})
	}
}

func deserializeTokenComponent(t *testing.T, serializedComponent string) map[string]interface{} {
	componentJSON, err := base64.RawURLEncoding.DecodeString(serializedComponent)
	require.NoError(t, err)
	var component map[string]interface{}
	err = json.Unmarshal([]byte(componentJSON), &component)
	require.NoError(t, err)
	return component
}
