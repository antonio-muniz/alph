package jwt_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	scenarios := []struct {
		description       string
		encodedToken      string
		signingAlgorithm  string
		signingKey        string
		expectedSignature string
	}{
		{
			description:       "signs_an_encoded_token_using_HMAC_SHA256",
			encodedToken:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhbHBoIiwic3ViIjoic29tZW9uZUBleGFtcGxlLm9yZyIsImF1ZCI6ImV4YW1wbGUub3JnIiwiaWF0IjoiMjAyMC0wNS0yNFQyMDowMDowMFoiLCJleHAiOiIyMDIwLTA1LTI0VDIwOjMwOjAwWiJ9",
			signingAlgorithm:  "HS256",
			signingKey:        "zLcwW6w2MEwS8RMzP71azVbQJyOK4fiV",
			expectedSignature: "AVC8mWAWEkQYYeduwnQVGyaOUXHKpQkbx4GT-iv7bOY",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			signedToken, err := jwt.Sign(
				scenario.encodedToken,
				scenario.signingAlgorithm,
				scenario.signingKey,
			)
			require.NoError(t, err)

			fmt.Println(signedToken)

			signedTokenComponents := strings.SplitN(signedToken, ".", 3)
			require.Len(t, signedTokenComponents, 3)

			signature := signedTokenComponents[2]
			require.Equal(t, scenario.expectedSignature, signature)
		})
	}
}
