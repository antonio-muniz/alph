package internalhelpers

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/stretchr/testify/require"
)

func InitializeSystem(t *testing.T, ctx context.Context) system.System {
	sys, err := internal.System(ctx)
	require.NoError(t, err)
	return sys
}
