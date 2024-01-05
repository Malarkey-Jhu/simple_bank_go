package token

import "time"

type Maker interface {
	// CreateToken creates a new token from specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifyToken checks if the token id valid or not
	VerifyToken(token string) (*Payload, error)
}
