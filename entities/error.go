package entities

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// ErrExampleInvalidRequest is used to show an example of a 400 Bad Request error
type ErrExampleInvalidRequest struct {
	Code    int          `json:"code" example:"2001"`
	Message string       `json:"message" example:"invalid request body"`
	Details []FieldError `json:"details,omitempty"`
}

type FieldError struct {
	Field   string `json:"field" example:"title"`
	Message string `json:"message" example:"Title is required"`
}

// ErrExampleIncorrectPassword is used to show an example of a 401 Unauthorized error
type ErrExampleIncorrectPassword struct {
	Code    int    `json:"code" example:"4002"`
	Message string `json:"message" example:"password is incorrect"`
}

// ErrExampleUserNotFound is used to show an example of a 404 Not Found error
type ErrExampleUserNotFound struct {
	Code    int    `json:"code" example:"4001"`
	Message string `json:"message" example:"user not found"`
}

// ErrExampleHashPassword is used to show an example of a 400 Bad Request error
type ErrExampleHashPassword struct {
	Code    int    `json:"code" example:"2007"`
	Message string `json:"message" example:"failed to hash password"`
}

// ErrExampleUserExists is used to show an example of a 409 Conflict error
type ErrExampleUserExists struct {
	Code    int    `json:"code" example:"4003"`
	Message string `json:"message" example:"user already exists"`
}

// ErrExampleInternalError is used to show an example of a 500 Internal Server Error
type ErrExampleInternalError struct {
	Code    int    `json:"code" example:"5000"`
	Message string `json:"message" example:"internal server error"`
}

type ErrExampleUnauthorized struct {
	Code    int    `json:"code" example:"1001"`
	Message string `json:"message" example:"unauthorized"`
}

// ErrExampleTaskNotFound is used to show an example of a 404 Not Found error
type ErrExampleTaskNotFound struct {
	Code    int    `json:"code" example:"3001"`
	Message string `json:"message" example:"task not found"`
}

// ErrExampleTaskAlreadyExists is used to show an example of a 409 Conflict error
type ErrExampleTaskAlreadyExists struct {
	Code    int    `json:"code" example:"3002"`
	Message string `json:"message" example:"task already exists"`
}
