package token

type Header struct {
	SignatureAlgorithm string `json:"alg"`
	TokenType          string `json:"typ"`
}
