package config

import "strings"

func LoadConfiguration() Config {
	return Config{
		JwtSignatureKey: "zLcwW6w2MEwS8RMzP71azVbQJyOK4fiV",
		JwtEncryptionPublicKey: strings.Join(
			[]string{
				"-----BEGIN PUBLIC KEY-----",
				"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzO8HOyr435l84SaPlOT0",
				"A5HCyL+DwL4RbTutK+t9osU0jSCNaJ39sK4OU54kI0blkEZMknVWdkHwb++oq7K2",
				"bt68tM5bD1E5wy+pLC07XfmXWNdtJQWbIEyfQCIsozqVoH407xqYT70FbJSCobf+",
				"TM/b9PUU3VbzK4qwvbsWgDRQToYUID9uJu0yg8hjFy2yeMX+J8gg6e/DsqlVXvca",
				"LhVdPT1+D0IYMzOPuNNYdMWvuPqRuN5Nyj2ckCPe7zJJvQE2ri2y5Oaac6a4otqP",
				"J4+laFTObq7N0EKj+Qr1ccBoIiYPHZo7l/ZfBoVpKBZwuVOkjW+WyDrg3ZL5o7f3",
				"1QIDAQAB",
				"-----END PUBLIC KEY-----",
			},
			"\n",
		),
	}
}
