package internalhelpers

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage/memory"
	"github.com/antonio-muniz/alph/pkg/clock"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/antonio-muniz/alph/test/helpers"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/require"
)

func InitializeSystem(t *testing.T, ctx context.Context) system.System {
	builder, err := di.NewBuilder()
	require.NoError(t, err)
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
		di.Def{
			Name: "clock",
			Build: func(container di.Container) (interface{}, error) {
				return clock.NewFrozenClock(helpers.Now()), nil
			},
		},
	)
	container := builder.Build()
	system := system.New(container)
	return system
}
