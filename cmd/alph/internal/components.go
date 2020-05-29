package internal

import (
	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

func Components() (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, errors.Wrap(err, "error building components")
	}
	builder.Add(di.Def{
		Name: "config",
		Build: func(container di.Container) (interface{}, error) {
			return config.LoadConfiguration(), nil
		},
	})
	container := builder.Build()
	return container, nil
}
