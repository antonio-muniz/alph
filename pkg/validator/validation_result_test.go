package validator_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/validator"
	"github.com/stretchr/testify/require"
)

func TestInvalid(t *testing.T) {
	scenarios := []struct {
		description string
		errors      []validator.Error
		expected    bool
	}{
		{
			description: "returns_false_when_there_is_no_validation_error",
			errors:      []validator.Error{},
			expected:    false,
		},
		{
			description: "returns_true_when_there_are_validation_errors",
			errors: []validator.Error{
				{
					Type:  "MISSING",
					Field: "MyRequiredField",
				},
			},
			expected: true,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := validator.Result{Errors: scenario.errors}
			require.Equal(t, scenario.expected, result.Invalid())
		})
	}
}
