package utils

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
)

// IPayload is the interface for the payload
type IPayloadConstruct interface {
	NewCreatePayload(userId string, duration time.Duration) (*Payload, error)
	GetAuthPayload(ctx context.Context, log *log.Logger) (*Payload, error)
	Valid(payload *Payload) error
}

type PayloadConstruct struct {
	config *config.Config
	log    *log.Logger
}

func NewPayloadConstruct(config *config.Config, log *log.Logger) IPayloadConstruct {
	return &PayloadConstruct{
		config: config,
		log:    log,
	}
}

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserId    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expires_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func (p PayloadConstruct) NewCreatePayload(userId string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserId:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// GetAuthPayload gets the auth payload from the context
func (p PayloadConstruct) GetAuthPayload(ctx context.Context, log *log.Logger) (*Payload, error) {
	log.DebugWithID(ctx, "[Utils: GetAuthPayload] Getting auth payload")
	raw := ctx.Value(constants.AuthorizationPayloadKey)
	payload, ok := raw.(*Payload)
	if !ok || payload == nil {
		log.ErrorWithID(ctx, "[Utils: GetAuthPayload] Unauthorized: token payload missing")
		return nil, constants.ErrUnauthorized
	}

	log.DebugWithID(ctx, "[Utils: GetAuthPayload] Auth payload found", payload)
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (p PayloadConstruct) Valid(payload *Payload) error {
	if time.Now().After(payload.ExpiredAt) {
		return constants.ErrExpiredToken
	}

	return nil
}
