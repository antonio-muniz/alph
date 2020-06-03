package helpers

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/stretchr/testify/require"
)

func HashPassword(t *testing.T, plainPassword string) string {
	hashedPassword, err := password.Hash(plainPassword)
	require.NoError(t, err)
	return hashedPassword
}
