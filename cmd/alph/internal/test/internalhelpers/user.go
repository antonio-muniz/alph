package internalhelpers

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/pkg/system"
	"github.com/stretchr/testify/require"
)

func CreateUser(
	t *testing.T,
	ctx context.Context,
	sys system.System,
	user model.User,
) {
	database := sys.Get("database").(storage.Database)
	err := database.CreateUser(ctx, user)
	require.NoError(t, err)
}
