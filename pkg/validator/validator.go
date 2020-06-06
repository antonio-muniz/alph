package validator

import (
	"fmt"
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

func (v Validator) Validate(payload interface{}) (Result, error) {
	err := ensurePayloadIsSupported(payload)
	if err != nil {
		return Result{}, err
	}
	err = v.validate.Struct(payload)
	switch typedErr := err.(type) {
	case nil:
		return Result{}, nil
	case validation.ValidationErrors:
		errors, err := convertValidationErrors(typedErr)
		if err != nil {
			return Result{}, err
		}
		return Result{Errors: errors}, nil
	default:
		return Result{}, errors.WithStack(err)
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

func convertValidationErrors(validationErrors validation.ValidationErrors) ([]Error, error) {
	var errors []Error
	for _, validationError := range validationErrors {
		error := convertValidationError(validationError)
		errors = append(errors, error)
	}
	return errors, nil
}

var convertionFunctions = map[string]func(validation.FieldError) Error{
	"required": func(validationError validation.FieldError) Error {
		return Error{
			Type:  "MISSING",
			Field: validationError.Field(),
		}
	},
	"gte": func(validationError validation.FieldError) Error {
		return Error{
			Type:  "TOO_LOW",
			Field: validationError.Field(),
			Value: validationError.Value(),
			Details: map[string]interface{}{
				"minimum": validationError.Param(),
			},
		}
	},
}

func convertValidationError(validationError validation.FieldError) Error {
	var error Error
	tag := validationError.ActualTag()
	switch tag {
	case "required":
		error = Error{
			Type:  "MISSING",
			Field: validationError.Field(),
		}
	case "gte":
		error = Error{
			Type:  "TOO_LOW",
			Field: validationError.Field(),
			Value: validationError.Value(),
			Details: map[string]interface{}{
				"minimum": validationError.Param(),
			},
		}
	default:
		panic(fmt.Sprintf("Unexpected validation tag '%s'", tag))
	}
	return error
}

func ensurePayloadIsSupported(payload interface{}) error {
	payloadType := reflect.TypeOf(payload)
	if payloadType.Kind() != reflect.Struct {
		return ErrNonStructPayload
	}
	fieldCount := payloadType.NumField()
	for fieldIndex := 0; fieldIndex < fieldCount; fieldIndex++ {
		field := payloadType.Field(fieldIndex)
		validateTagValue := field.Tag.Get("validate")
		if validateTagValue == "" {
			continue
		}
		validationRules := strings.Split(validateTagValue, ",")
		for _, rule := range validationRules {
			err := ensureRuleIsSupported(rule)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ensureRuleIsSupported(rule string) error {
	ruleParts := strings.SplitN(rule, "=", 2)
	tag := ruleParts[0]
	for supportedTag := range convertionFunctions {
		if tag == supportedTag {
			return nil
		}
	}
	return ErrUnsupportedValidationTag{Tag: tag}
}
