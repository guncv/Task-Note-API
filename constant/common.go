package constants

type TaskStatus string

const (
	TaskStatusPending       TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted     TaskStatus = "COMPLETED"
	Alphabet                           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AuthorizationHeaderKey             = "authorization"
	AuthorizationTypeBearer            = "bearer"
	AuthorizationPayloadKey            = "authorization_payload"
	CurrentTimeLocation                = "Asia/Bangkok"
	TestAppEnv                         = "test"
)
