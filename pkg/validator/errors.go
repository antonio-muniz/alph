package validator

import (
	"errors"
	"fmt"
)

var ErrNonStructPayload = errors.New("validation payload is not a struct")

type ErrUnsupportedValidationTag struct {
	Tag string
}

func (e ErrUnsupportedValidationTag) Error() string {
	return fmt.Sprintf("unsupported validation tag '%s'", e.Tag)
}
