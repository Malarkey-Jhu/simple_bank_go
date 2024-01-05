package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

var ErrExpiredToken = errors.New("token is expired")

type PasteoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasteoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasteoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasteoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasteoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	if time.Now().After(payload.ExpiresAt.Time) {
		return nil, ErrExpiredToken
	}

	return payload, nil
}
