package validator_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/validator"
	"github.com/stretchr/testify/require"
)

type payload struct {
	Name                      string `json:"name" validate:"required"`
	HoursWastedReadingBadCode int    `json:"hours_wasted_reading_bad_code" validate:"gte=0"`
}

type payloadWithoutRules struct {
	Name                      string `json:"name"`
	HoursWastedReadingBadCode int    `json:"hours_wasted_reading_bad_code"`
}

type payloadWithUnsupportedTag struct {
	Name                      string `json:"name" validate:"required,dont_do_this"`
	HoursWastedReadingBadCode int    `json:"hours_wasted_reading_bad_code" validate:"gte=0"`
}

func TestValidate(t *testing.T) {
	scenarios := []struct {
		description      string
		validatorOptions []validator.Option
		payload          interface{}
		expectedResult   validator.Result
		expectedError    error
	}{
		{
			description: "returns_no_error_for_valid_payload",
			payload: payload{
				Name:                      "Someone",
				HoursWastedReadingBadCode: 99999,
			},
			expectedResult: validator.Result{},
		},
		{
			description: "returns_no_error_for_payload_without_rules",
			payload: payloadWithoutRules{
				Name:                      "",
				HoursWastedReadingBadCode: -1,
			},
			expectedResult: validator.Result{},
		},
		{
			description: "returns_validation_error_for_invalid_payload",
			payload: payload{
				Name:                      "",
				HoursWastedReadingBadCode: 99999,
			},
			expectedResult: validator.Result{
				Errors: []validator.Error{
					{
						Type:  "MISSING",
						Field: "Name",
					},
				},
			},
		},
		{
			description: "returns_validation_error_for_invalid_payload_with_json_field_name",
			validatorOptions: []validator.Option{
				validator.ErrorFieldFromJSONTag(),
			},
			payload: payload{
				Name:                      "Someone",
				HoursWastedReadingBadCode: -1,
			},
			expectedResult: validator.Result{
				Errors: []validator.Error{
					{
						Type:  "TOO_LOW",
						Field: "hours_wasted_reading_bad_code",
						Value: -1,
						Details: map[string]interface{}{
							"minimum": "0",
						},
					},
				},
			},
		},
		{
			description: "returns_multiple_validation_errors_when_needed",
			payload: payload{
				Name:                      "",
				HoursWastedReadingBadCode: -1,
			},
			expectedResult: validator.Result{
				Errors: []validator.Error{
					{
						Type:  "MISSING",
						Field: "Name",
					},
					{
						Type:  "TOO_LOW",
						Field: "HoursWastedReadingBadCode",
						Value: -1,
						Details: map[string]interface{}{
							"minimum": "0",
						},
					},
				},
			},
		},
		{
			description: "fails_when_there_is_an_unsupported_tag_in_a_struct_field",
			payload: payloadWithUnsupportedTag{
				Name:                      "Someone",
				HoursWastedReadingBadCode: 99999,
			},
			expectedResult: validator.Result{},
			expectedError:  validator.ErrUnsupportedValidationTag{Tag: "dont_do_this"},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			validator := validator.New(scenario.validatorOptions...)
			result, err := validator.Validate(scenario.payload)
			require.Equal(t, scenario.expectedResult, result)
			require.Equal(t, scenario.expectedError, err)
		})
	}
}
