package encryption_test

import (
	"testing"

	"github.com/antonio-muniz/alph/pkg/encryption"
	"github.com/stretchr/testify/require"
)

func TestAESEncrypt(t *testing.T) {
	scenarios := []struct {
		description   string
		message       string
		encryptionKey string
		expectedError error
	}{
		{
			description:   "encryption_fails_when_the_key_is_too_short",
			message:       "shhh-this-is-a-huge-secret",
			encryptionKey: "hard-to-guess",
			expectedError: encryption.ErrBadAESKeyLength,
		},
		{
			description:   "decryption_fails_when_the_key_is_too_long",
			message:       "shhh-this-is-a-huge-secret",
			encryptionKey: "dont-share-this-key-with-someone-even-with-friends-or-family",
			expectedError: encryption.ErrBadAESKeyLength,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			_, err := encryption.AESEncrypt(scenario.message, scenario.encryptionKey)
			if scenario.expectedError != nil {
				require.Equal(t, scenario.expectedError, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestAESDecrypt(t *testing.T) {
	scenarios := []struct {
		description      string
		encryptedMessage string
		decryptionKey    string
		expectedError    error
	}{
		{
			description:      "decryption_fails_when_the_key_is_too_short",
			encryptedMessage: "VERYENCRYPTEDMESSAGE",
			decryptionKey:    "hard-to-guess",
			expectedError:    encryption.ErrBadAESKeyLength,
		},
		{
			description:      "decryption_fails_when_the_key_is_too_long",
			encryptedMessage: "VERYENCRYPTEDMESSAGE",
			decryptionKey:    "dont-share-this-key-with-someone-even-with-friends-or-family",
			expectedError:    encryption.ErrBadAESKeyLength,
		},
		{
			description:      "decryption_fails_when_the_message_is_not_in_base_64",
			encryptedMessage: "VERYENCRYPTEDMESSAGE ;)",
			decryptionKey:    "dont-share-this-key-with-anybody",
			expectedError:    encryption.ErrBadEncryptedMessageEncoding,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			_, err := encryption.AESDecrypt(scenario.encryptedMessage, scenario.decryptionKey)
			if scenario.expectedError != nil {
				require.Equal(t, scenario.expectedError, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestAESEncryptDecrypt(t *testing.T) {
	scenarios := []struct {
		description   string
		message       string
		encryptionKey string
		decryptionKey string
		expectedMatch bool
	}{
		{
			description:   "encrypts_then_decrypts_a_message_using_AES_algorithm",
			message:       "shhh-this-is-a-huge-secret",
			encryptionKey: "dont-share-this-key-with-anybody",
			decryptionKey: "dont-share-this-key-with-anybody",
			expectedMatch: true,
		},
		{
			description:   "message_does_not_match_when_the_decryption_key_is_incorrect",
			message:       "shhh-this-is-a-huge-secret",
			encryptionKey: "dont-share-this-key-with-anybody",
			decryptionKey: "dont-share-this-key-with-someone",
			expectedMatch: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			encryptedMessage, err := encryption.AESEncrypt(scenario.message, scenario.encryptionKey)
			require.NoError(t, err)
			decryptedMessage, err := encryption.AESDecrypt(encryptedMessage, scenario.decryptionKey)
			require.NoError(t, err)
			if scenario.expectedMatch {
				require.Equal(t, scenario.message, decryptedMessage)
			} else {
				require.NotEqual(t, scenario.message, decryptedMessage)
			}
		})
	}
}
