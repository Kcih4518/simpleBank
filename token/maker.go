package token

import "time"

// Maker is a generic interface to generate tokens
type Maker interface {
	// CreateToken generates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
