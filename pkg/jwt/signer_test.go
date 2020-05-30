package jwt_test

import (
	"strings"
	"testing"

	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	scenarios := []struct {
		description       string
		encodedToken      string
		signingKey        string
		expectedSignature string
	}{
		{
			description:       "signs_an_encoded_token_using_HMAC_SHA256",
			encodedToken:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhbHBoIiwic3ViIjoic29tZW9uZUBleGFtcGxlLm9yZyIsImF1ZCI6ImV4YW1wbGUub3JnIiwiaWF0IjoiMjAyMC0wNS0yNFQyMDowMDowMFoiLCJleHAiOiIyMDIwLTA1LTI0VDIwOjMwOjAwWiJ9",
			signingKey:        "zLcwW6w2MEwS8RMzP71azVbQJyOK4fiV",
			expectedSignature: "AVC8mWAWEkQYYeduwnQVGyaOUXHKpQkbx4GT-iv7bOY",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			signedToken, err := jwt.Sign(scenario.encodedToken, scenario.signingKey)
			require.NoError(t, err)

			signedTokenComponents := strings.SplitN(signedToken, ".", 3)
			require.Len(t, signedTokenComponents, 3)

			signature := signedTokenComponents[2]
			require.Equal(t, scenario.expectedSignature, signature)
		})
	}
}

func TestCheckSignature(t *testing.T) {
	scenarios := []struct {
		description   string
		signedToken   string
		signingKey    string
		expectedError error
	}{
		{
			description:   "returns_true_if_the_signature_is_correct",
			signedToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhbHBoIiwic3ViIjoic29tZW9uZUBleGFtcGxlLm9yZyIsImF1ZCI6ImV4YW1wbGUub3JnIiwiaWF0IjoiMjAyMC0wNS0yNFQyMDowMDowMFoiLCJleHAiOiIyMDIwLTA1LTI0VDIwOjMwOjAwWiJ9.AVC8mWAWEkQYYeduwnQVGyaOUXHKpQkbx4GT-iv7bOY",
			signingKey:    "zLcwW6w2MEwS8RMzP71azVbQJyOK4fiV",
			expectedError: nil,
		},
		{
			description:   "returns_false_if_the_signature_is_incorrect",
			signedToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhbHBoIiwic3ViIjoic29tZW9uZUBleGFtcGxlLm9yZyIsImF1ZCI6ImV4YW1wbGUub3JnIiwiaWF0IjoiMjAyMC0wNS0yNFQyMDowMDowMFoiLCJleHAiOiIyMDIwLTA1LTI0VDIwOjMwOjAwWiJ9.BVC8mWAWEkQYYeduwnQVGyaOUXHKpQkbx4GT-iv7bOY",
			signingKey:    "zLcwW6w2MEwS8RMzP71azVbQJyOK4fiV",
			expectedError: jwt.ErrIncorrectSignature,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			err := jwt.CheckSignature(scenario.signedToken, scenario.signingKey)
			require.Equal(t, scenario.expectedError, err)
		})
	}
}
