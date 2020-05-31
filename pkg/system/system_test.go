package system_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/require"
)

func TestSystem(t *testing.T) {
	containerBuilder, err := di.NewBuilder()
	require.NoError(t, err)
	containerBuilder.Add(di.Def{
		Name: "useful-thing",
		Build: func(container di.Container) (interface{}, error) {
			return "this-is-useful", nil
		},
	})
	container := containerBuilder.Build()
	system := system.New(container)
	component := system.Get("useful-thing")
	require.Equal(t, "this-is-useful", component)
}
