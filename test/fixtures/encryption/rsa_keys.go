package encryption

import (
	"strings"
)

var publicKey = strings.Join(
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
)

var privateKey = strings.Join(
	[]string{
		"-----BEGIN PRIVATE KEY-----",
		"MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDM7wc7KvjfmXzh",
		"Jo+U5PQDkcLIv4PAvhFtO60r632ixTSNII1onf2wrg5TniQjRuWQRkySdVZ2QfBv",
		"76irsrZu3ry0zlsPUTnDL6ksLTtd+ZdY120lBZsgTJ9AIiyjOpWgfjTvGphPvQVs",
		"lIKht/5Mz9v09RTdVvMrirC9uxaANFBOhhQgP24m7TKDyGMXLbJ4xf4nyCDp78Oy",
		"qVVe9xouFV09PX4PQhgzM4+401h0xa+4+pG43k3KPZyQI97vMkm9ATauLbLk5ppz",
		"prii2o8nj6VoVM5urs3QQqP5CvVxwGgiJg8dmjuX9l8GhWkoFnC5U6SNb5bIOuDd",
		"kvmjt/fVAgMBAAECggEAGVNUvnAiD2flceGVDt54dVR3EN3yB0i12JzuWqYggZQD",
		"WYlzUEFuD36DELxTVPS2++xkHBlaFQUzFHI2kvlj6DGoemOiBzOPgtqJ+oagdo2Y",
		"sYb12wRlVkmByKgwgf+EbDAMlJvhxMDkQbXcquWVDKMHWK9M48mSBYh/LQNI1ZDo",
		"tTKrypBjADClGtZngYTkyv4rl0LLLNYum/gn7kKAWDxW9UodHHAT63KQ1PMRHu2M",
		"nNNIofLgkXnaKWKZTxvDC4IxHi1tGjNb3hBfIcXC6CrvUHN0Y+EpGoI8MNUxXziM",
		"CpLAyZCc+BdI9HzwBQZmaYq7osi75axnCZCbw/1s4QKBgQD2N30iuLqXvOXMwGAX",
		"ZVfN/KdkXF1DEpPSCQ8LPfan5EX9FPmlQC9uGnsAczDO14lr8hxvle6Dr0L29O9r",
		"Jvoh+kO0gfGkfvoTJwQsbtFeZLTl/AQDCrzYkH0wl1epAfd94SwcxKiaDQmQPaLy",
		"B1rzcC/ptJ2+FnL+pWCwsrA3XQKBgQDVE5v0v1qvZAC3/ogNPO7yzqbA0WEeBNU+",
		"oBByUnzO4QFPu38p/ptMVAD3eFa0qmUmRiDQQyFFx5To1K1lubVqksAk0fliZQY5",
		"bvRgc7soi4tElDaYhW7gbejUm7ja8qXiWQmEt8oSTP7Krsd0kZgjC1b2UYWt4hnq",
		"8dCSUX6S2QKBgBCPeuMy+ZnrIqm65rusHVPFgpzFeaBhR4ABC/n9mPGB8RMsrf3n",
		"a0lEjrFhDg9bf+q6xh7bPsesqhxiLRhP7tMKOjOR6ebeg4N2RFmYuoxWybQOV1a9",
		"Ciam0UTLaSH89+CHnKfyskRqTBJfku4kgIGPMinN9C4s5F8sUGySGDghAoGBAJe6",
		"mvexu5XhlFFxQtiKZ20nkr8DWjDSKDesK4n9CzsWJqTsUyVIVLYJq5aNbtC/9b9H",
		"eN5Ur0DNWlBiAgfKMlFtpr2ReBCE9+W2CPcV8lRw6f8vu6Dog4f1PDMJoJzo5Q9F",
		"b+2+InrdsSILS6TKeJeb35UPgC3OL76J64gF2CghAoGAVfhnnlJs1+reWHOmcQKT",
		"FjzOCS7t0xyCvyVcnn0RnZbU6+sSl8yBFe6f54krley/+MkGcNAt2iFakE7Hdkw2",
		"Rujbz0vrXVeyKBOp4YMkFH+waicbPkQXdnSZqU0X9KeaKNIZFDSyozYu3gir04mN",
		"9cEGoXCdXHmIP1idUsHdXwA=",
		"-----END PRIVATE KEY-----",
	},
	"\n",
)

func PublicKey() string {
	return publicKey
}

func PrivateKey() string {
	return privateKey
}
