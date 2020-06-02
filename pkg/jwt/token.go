package jwt

type Token struct {
	Issuer         string    `json:"iss"`
	Subject        string    `json:"sub"`
	Audience       string    `json:"aud"`
	ExpirationTime Timestamp `json:"exp"`
}
