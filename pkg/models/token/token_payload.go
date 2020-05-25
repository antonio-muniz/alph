package token

import "time"

type TokenPayload struct {
	Issuer         string
	Subject        string
	Audience       string
	IssuedAt       time.Time
	ExpirationTime time.Time
}
