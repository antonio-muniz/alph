package internalhelpers

import (
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/config"
	"github.com/antonio-muniz/alph/pkg/jwt"
	"github.com/antonio-muniz/alph/pkg/system"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/encryption"
	"github.com/stretchr/testify/require"
)

func VerifyAccessToken(
	t *testing.T,
	sys system.System,
	expectedToken jwt.Token,
	accessToken string,
) {
	config := sys.Get("config").(config.Config)
	unpackedToken, err := jwt.Unpack(
		accessToken,
		jwt.UnpackSettings{
			DecryptionKey: fixtures.PrivateKey(),
			SignatureKey:  config.JWTSignatureKey,
		},
	)
	require.NoError(t, err)
	require.Equal(t, expectedToken, unpackedToken)
}
