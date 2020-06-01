package validator

import (
	"fmt"
	"strings"
)

type Error struct {
	Code  string
	Field string
	Value interface{}
}

type Errors []Error

func (e Error) Error() string {
	return fmt.Sprintf("'%s'='%v': %s", e.Field, e.Value, e.Code)
}

func (e Errors) Error() string {
	var errorMessages []string
	for _, err := range e {
		errorMessages = append(errorMessages, err.Error())
	}
	return strings.Join(errorMessages, ";")
}
