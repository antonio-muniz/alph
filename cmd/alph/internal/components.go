package internal

import (
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

func initializeComponents() (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, errors.Wrap(err, "error building components")
	}
	container := builder.Build()
	return container, nil
}
