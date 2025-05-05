package utils

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/o1egl/paseto"
)

type IPasetoMaker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(config *config.Config) (IPasetoMaker, error) {
	if len(config.TokenConfig.TokenSymmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: mush be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(config.TokenConfig.TokenSymmetricKey),
	}
	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	if err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil); err != nil {
		return nil, constants.ErrInvalidToken
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
