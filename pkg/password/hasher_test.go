package password_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"
)

func TestHashValidate(t *testing.T) {
	scenarios := []struct {
		description              string
		password                 string
		attemptedPassword        string
		expectedValidationResult bool
	}{
		{
			description:              "hashes_password_and_matches_the_same_password",
			password:                 "myreallystrongpassword",
			attemptedPassword:        "myreallystrongpassword",
			expectedValidationResult: true,
		},
		{
			description:              "hashes_password_and_does_not_match_a_different_password",
			password:                 "myreallystrongpassword",
			attemptedPassword:        "myreallywrongpassword",
			expectedValidationResult: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			hashedPassword, err := password.Hash(scenario.password)
			require.NoError(t, err)
			otherHashedPassword, err := password.Hash(scenario.password)
			require.NoError(t, err)
			require.NotEqual(t, hashedPassword, otherHashedPassword)
			validationResult, err := password.Validate(scenario.attemptedPassword, hashedPassword)
			require.NoError(t, err)
			require.Equal(t, scenario.expectedValidationResult, validationResult)
			otherValidationResult, err := password.Validate(scenario.attemptedPassword, otherHashedPassword)
			require.NoError(t, err)
			require.Equal(t, scenario.expectedValidationResult, otherValidationResult)
		})
	}
}
