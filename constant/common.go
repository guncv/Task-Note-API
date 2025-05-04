package constants

import "errors"

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	Alphabet                       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)
