package constants

import "errors"

type TaskStatus string

const (
	TaskStatusPending       TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted     TaskStatus = "COMPLETED"
	Alphabet                           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AuthorizationHeaderKey             = "authorization"
	AuthorizationTypeBearer            = "bearer"
	AuthorizationPayloadKey            = "authorization_payload"
	CurrentTimeLocation                = "Asia/Bangkok"
)

var (
	// Verify Token
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")

	// Authorization Middleware
	ErrAuthorizationHeaderNotProvided         = errors.New("authorization header is not provided")
	ErrInvalidAuthorizationHeaderFormat       = errors.New("invalid authorization header format")
	ErrAuthorizationHeaderMustStartWithBearer = errors.New("authorization header must start with " + AuthorizationTypeBearer)
	ErrFailedToVerifyToken                    = errors.New("failed to verify token")
	ErrUserIdDoesNotMatchWithYourAccount      = errors.New("user id of this task does not match with your account")
)
