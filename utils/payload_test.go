package utils

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/stretchr/testify/require"
)

func TestPayload_GetAuthPayload(t *testing.T) {
	log := log.Initialize(constants.TestAppEnv)

	// Valid payload in context
	expectedPayload := &Payload{
		ID:        uuid.New(),
		UserId:    "1",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	config := &config.Config{
		TokenConfig: config.TokenConfig{
			TokenSymmetricKey:   RandomString(32),
			AccessTokenDuration: time.Minute,
		},
	}
	// Set context with correct payload
	ctx := context.WithValue(context.Background(), constants.AuthorizationPayloadKey, expectedPayload)

	gotPayload, err := NewPayloadConstruct(config, log).GetAuthPayload(ctx, log)
	require.NoError(t, err)
	require.Equal(t, expectedPayload.ID, gotPayload.ID)
	require.Equal(t, expectedPayload.UserId, gotPayload.UserId)
	require.WithinDuration(t, expectedPayload.IssuedAt, gotPayload.IssuedAt, time.Second)
	require.WithinDuration(t, expectedPayload.ExpiredAt, gotPayload.ExpiredAt, time.Second)

	// Context has no payload
	ctx = context.Background()

	_, err = NewPayloadConstruct(config, log).GetAuthPayload(ctx, log)
	require.Error(t, err)
	require.EqualError(t, err, "unauthorized: token payload missing")
}
