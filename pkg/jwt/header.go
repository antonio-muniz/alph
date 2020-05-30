package jwt

type Header struct {
	SignatureAlgorithm string `json:"alg"`
	TokenType          string `json:"typ"`
}
