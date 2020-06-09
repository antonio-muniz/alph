package internal

import (
	"context"
	"os"

	"github.com/antonio-muniz/alph/pkg/clock"
	"github.com/antonio-muniz/alph/pkg/logger"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage/memory"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

func System(_ctx context.Context) (system.System, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return system.System{}, errors.WithStack(err)
	}
	builder.Add(
		di.Def{
			Name: "config",
			Build: func(container di.Container) (interface{}, error) {
				return config.Load(), nil
			},
		},
		di.Def{
			Name: "database",
			Build: func(container di.Container) (interface{}, error) {
				return memory.NewDatabase(), nil
			},
		},
		di.Def{
			Name: "clock",
			Build: func(container di.Container) (interface{}, error) {
				return clock.NewWorkingClock(), nil
			},
		},
		di.Def{
			Name: "logger",
			Build: func(container di.Container) (interface{}, error) {
				return logger.NewLogger(os.Stdout), nil
			},
		},
	)
	container := builder.Build()
	system := system.New(container)
	return system, nil
}
