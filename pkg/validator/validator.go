package validator

import (
	"reflect"
	"strings"

	validation "github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type Validator struct {
	validate *validation.Validate
}

type Option func(*Validator)

func New(options ...Option) Validator {
	validate := validation.New()
	validator := Validator{validate: validate}
	for _, option := range options {
		option(&validator)
	}
	return validator
}

func (v Validator) Validate(payload interface{}) error {
	err := v.validate.Struct(payload)
	switch typedErr := err.(type) {
	case nil:
		return nil
	case validation.ValidationErrors:
		return convertValidationErrors(typedErr)
	default:
		return errors.Wrap(err, "validating payload")
	}
}

func ErrorFieldFromJSONTag() Option {
	return func(validator *Validator) {
		validator.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			jsonTagValue := field.Tag.Get("json")
			jsonTagParts := strings.SplitN(jsonTagValue, ",", 2)
			jsonName := jsonTagParts[0]
			return jsonName
		})
	}
}

func convertValidationErrors(validationErrors validation.ValidationErrors) Errors {
	var errors Errors
	for _, validationError := range validationErrors {
		errors = append(errors, Error{
			Code:  errorCodeFromValidationTag(validationError.Tag()),
			Field: validationError.Field(),
			Value: validationError.Value(),
		})
	}
	return errors
}

func errorCodeFromValidationTag(tag string) string {
	switch tag {
	case "gte":
		return "TOO_LOW"
	default:
		return ""
	}
}
