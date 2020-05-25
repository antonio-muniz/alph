package token

type Payload struct {
	Issuer         string    `json:"iss"`
	Subject        string    `json:"sub"`
	Audience       string    `json:"aud"`
	IssuedAt       Timestamp `json:"iat"`
	ExpirationTime Timestamp `json:"exp"`
}
