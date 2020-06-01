package validator_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/validator"
	"github.com/stretchr/testify/require"
)

type samplePayload struct {
	Name                      string
	HoursWastedReadingBadCode int `json:"hours_wasted_reading_bad_code" validate:"gte=0"`
}

func TestValidate(t *testing.T) {
	scenarios := []struct {
		description      string
		validatorOptions []validator.Option
		payload          samplePayload
		expectedError    error
	}{
		{
			description: "returns_no_error_for_valid_payload",
			payload: samplePayload{
				Name:                      "Someone",
				HoursWastedReadingBadCode: 99999,
			},
			expectedError: nil,
		},
		{
			description: "returns_error_for_invalid_payload",
			payload: samplePayload{
				Name:                      "Someone",
				HoursWastedReadingBadCode: -1,
			},
			expectedError: validator.Errors(
				[]validator.Error{
					{
						Code:  "TOO_LOW",
						Field: "HoursWastedReadingBadCode",
						Value: -1,
					},
				},
			),
		},
		{
			description: "returns_error_for_invalid_payload_with_json_field_name",
			validatorOptions: []validator.Option{
				validator.ErrorFieldFromJSONTag(),
			},
			payload: samplePayload{
				Name:                      "Someone",
				HoursWastedReadingBadCode: -1,
			},
			expectedError: validator.Errors(
				[]validator.Error{
					{
						Code:  "TOO_LOW",
						Field: "hours_wasted_reading_bad_code",
						Value: -1,
					},
				},
			),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			validator := validator.New(scenario.validatorOptions...)
			err := validator.Validate(scenario.payload)
			require.Equal(t, scenario.expectedError, err)
		})
	}
}
