package internalhelpers

import (
	"context"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/pkg/password"
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

func VerifyUserExists(
	t *testing.T,
	ctx context.Context,
	sys system.System,
	expectedUsername string,
	expectedPassword string,
) {
	database := sys.Get("database").(storage.Database)
	user, err := database.GetUser(ctx, expectedUsername)
	require.NoError(t, err)
	passwordMatch, err := password.Validate(expectedPassword, user.HashedPassword)
	require.NoError(t, err)
	require.True(t, passwordMatch)
}

func VerifyUserDoesNotExist(
	t *testing.T,
	ctx context.Context,
	sys system.System,
	username string,
) {
	database := sys.Get("database").(storage.Database)
	_, err := database.GetUser(ctx, username)
	require.Equal(t, storage.ErrUserNotFound, err)
}
