package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTMaker is a JSON Web Token maker
// It will implement by symetric encryption
type JWTMaker struct {
	secretKey string
}

// minSecretKeySize is the minimum key size in bytes
const minSecretKeySize = 32

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	// FIXME: need to implement Maker interface's methods
	//  "*JWTMaker does not implement Maker (missing method CreateToken)"
	return &JWTMaker{secretKey}, nil
}

// FIXME: Why need to using JWTmaker instead of Maker interface?
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error)
