package token

import "time"

type Payload struct {
	Issuer         string    `json:"iss"`
	Subject        string    `json:"sub"`
	Audience       string    `json:"aud"`
	IssuedAt       time.Time `json:"iat"`
	ExpirationTime time.Time `json:"exp"`
}
