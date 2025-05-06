package utils

import (
	"context"
	"fmt"

	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(ctx context.Context, password string, log *log.Logger) (string, error) {
	log.DebugWithID(ctx, "[Utils: HashPassword] Hashing password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.ErrorWithID(ctx, "[Utils: HashPassword] Failed to hash password", err)
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password
func CheckPassword(ctx context.Context, password string, hashedPassword string, log *log.Logger) error {
	log.DebugWithID(ctx, "[Utils: CheckPassword] Checking password")

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.ErrorWithID(ctx, "[Utils: CheckPassword] Failed to check password", err)
		return fmt.Errorf("failed to check password: %w", err)
	}
	return nil
}
