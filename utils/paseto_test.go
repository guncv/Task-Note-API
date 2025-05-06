package utils

import (
	"testing"
	"time"

	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	config := &config.Config{
		TokenConfig: config.TokenConfig{
			TokenSymmetricKey:   RandomString(32),
			AccessTokenDuration: time.Minute,
		},
	}

	log := log.Initialize(constants.TestAppEnv)
	maker, err := NewPasetoMaker(config, NewPayloadConstruct(config, log))

	require.NoError(t, err)

	userId := RandomString(32)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(userId, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userId, payload.UserId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestInvalidPasetoToken(t *testing.T) {
	config := &config.Config{
		TokenConfig: config.TokenConfig{
			TokenSymmetricKey:   RandomString(31),
			AccessTokenDuration: time.Minute,
		},
	}

	log := log.Initialize(constants.TestAppEnv)
	maker, err := NewPasetoMaker(config, NewPayloadConstruct(config, log))
	require.Error(t, err)
	require.Nil(t, maker)
}

func TestExpiredPasetoToken(t *testing.T) {
	config := &config.Config{
		TokenConfig: config.TokenConfig{
			TokenSymmetricKey:   RandomString(32),
			AccessTokenDuration: time.Minute,
		},
	}

	log := log.Initialize(constants.TestAppEnv)
	maker, err := NewPasetoMaker(config, NewPayloadConstruct(config, log))
	require.NoError(t, err)

	userId := RandomString(32)
	token, err := maker.CreateToken(userId, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, constants.ErrExpiredToken.Error())
	require.Nil(t, payload)
}
