package internal

import (
	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage/memory"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

func System() (system.System, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return system.System{}, errors.Wrap(err, "error building components")
	}
	builder.Add(
		di.Def{
			Name: "config",
			Build: func(container di.Container) (interface{}, error) {
				return config.LoadConfiguration(), nil
			},
		},
		di.Def{
			Name: "database",
			Build: func(container di.Container) (interface{}, error) {
				return memory.NewDatabase(), nil
			},
		},
	)
	container := builder.Build()
	system := system.New(container)
	return system, nil
}
