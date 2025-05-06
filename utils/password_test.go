package utils

import (
	"context"
	"testing"

	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)
	ctx := context.Background()

	log := log.Initialize(constants.TestAppEnv)
	hashedPassword1, err := HashPassword(ctx, password, log)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(ctx, password, hashedPassword1, log)
	require.NoError(t, err)

	wrongPassword := RandomString(8)
	err = CheckPassword(ctx, wrongPassword, hashedPassword1, log)
	require.ErrorContains(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(ctx, password, log)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
