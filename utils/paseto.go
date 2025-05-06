package utils

import (
	"errors"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/o1egl/paseto"
)

type IPasetoMaker interface {
	CreateToken(userId string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto           *paseto.V2
	symmetricKey     []byte
	payloadConstruct IPayloadConstruct
}

func NewPasetoMaker(config *config.Config, payloadConstruct IPayloadConstruct) (IPasetoMaker, error) {
	if len(config.TokenConfig.TokenSymmetricKey) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key size: must be exactly 32 bytes")
	}

	maker := &PasetoMaker{
		paseto:           paseto.NewV2(),
		symmetricKey:     []byte(config.TokenConfig.TokenSymmetricKey),
		payloadConstruct: payloadConstruct,
	}
	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(userId string, duration time.Duration) (string, error) {
	payload, err := maker.payloadConstruct.NewCreatePayload(userId, duration)
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

	if err := maker.payloadConstruct.Valid(payload); err != nil {
		return nil, err
	}

	return payload, nil
}
