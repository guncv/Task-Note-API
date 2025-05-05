package utils

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
)

// IPayload is the interface for the payload
type IPayload interface {
	Valid() error
	GetAuthPayload(ctx context.Context, log *log.Logger) (*Payload, error)
}

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserId    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expires_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(userId string, duration time.Duration) (IPayload, error) {
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
func (payload *Payload) GetAuthPayload(ctx context.Context, log *log.Logger) (*Payload, error) {
	log.DebugWithID(ctx, "[Utils: GetAuthPayload] Getting auth payload")
	raw := ctx.Value(constants.AuthorizationPayloadKey)
	payload, ok := raw.(*Payload)
	if !ok || payload == nil {
		log.ErrorWithID(ctx, "[Utils: GetAuthPayload] Unauthorized: token payload missing")
		return nil, errors.New("unauthorized: token payload missing")
	}

	log.DebugWithID(ctx, "[Utils: GetAuthPayload] Auth payload found", payload)
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return constants.ErrExpiredToken
	}

	return nil
}
