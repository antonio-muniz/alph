package encryption_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/encryption"
	"github.com/stretchr/testify/require"
)

func TestAESEncryptDecrypt(t *testing.T) {
	scenarios := []struct {
		description   string
		message       string
		encryptionKey string
	}{
		{
			description:   "encrypts_then_decrypts_a_message_using_AES_algorithm",
			message:       "shhh-this-is-a-huge-secret",
			encryptionKey: "dont-share-this-key-with-anybody",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			encryptedMessage, err := encryption.AESEncrypt(scenario.message, scenario.encryptionKey)
			require.NoError(t, err)
			decryptedMessage, err := encryption.AESDecrypt(encryptedMessage, scenario.encryptionKey)
			require.NoError(t, err)
			require.Equal(t, scenario.message, decryptedMessage)
		})
	}
}
