package constants

import (
	"errors"
	"net/http"
)

type ErrorType int

const (
	// Token-related errors
	CodeTokenExpired            ErrorType = 1001
	CodeTokenInvalid            ErrorType = 1002
	CodeTokenVerificationFailed ErrorType = 1003
	CodeAuthHeaderMissing       ErrorType = 1004
	CodeAuthHeaderFormatInvalid ErrorType = 1005
	CodeAuthHeaderMissingBearer ErrorType = 1006
	CodeUserIdMismatch          ErrorType = 1007
	CodeUnauthorized            ErrorType = 1008

	// Input & validation
	CodeInvalidRequestBody                ErrorType = 2001
	CodeMissingRequiredFields             ErrorType = 2002
	CodeInvalidFieldFormat                ErrorType = 2003
	CodeInvalidStatusTransition           ErrorType = 2004
	CodeInvalidQueryRequestParam          ErrorType = 2005
	CodeInvalidRequestParam               ErrorType = 2006
	CodeHashPassword                      ErrorType = 2007
	CodeConvertFileHeaderToBase64         ErrorType = 2008
	CodeOpenFileContext                   ErrorType = 2009
	CodeUserIdDoesNotMatchWithYourAccount ErrorType = 2010
	CodeOtherError                        ErrorType = 2012

	// Task Resource
	CodeTaskNotFound      ErrorType = 3001
	CodeTaskAlreadyExists ErrorType = 3002

	// User Resource
	CodeUserNotFound      ErrorType = 4001
	CodePasswordIncorrect ErrorType = 4002
	CodeUserAlreadyExists ErrorType = 4003

	// Internal
	CodeInternalServerError       ErrorType = 5000
	CodeServiceUnavailable        ErrorType = 5001
	CodeGetCurrentTimeWithRFC3339 ErrorType = 5002
)

var (
	// Token-related errors
	ErrExpiredToken            = errors.New("token has expired")                                     // 1001
	ErrInvalidToken            = errors.New("token is invalid")                                      // 1002
	ErrFailedToVerifyToken     = errors.New("failed to verify token")                                // 1003
	ErrAuthHeaderMissing       = errors.New("authorization header is not provided")                  // 1004
	ErrAuthHeaderFormatInvalid = errors.New("invalid authorization header format")                   // 1005
	ErrAuthHeaderMissingBearer = errors.New("authorization header must start with Bearer")           // 1006
	ErrUserIdMismatch          = errors.New("user id of this task does not match with your account") // 1007
	ErrUnauthorized            = errors.New("unauthorized: token payload is invalid")                // 1008

	// Input & validation
	ErrInvalidRequestBody                = errors.New("invalid request body")                                  // 2001
	ErrMissingRequiredFields             = errors.New("missing required fields")                               // 2002
	ErrInvalidFieldFormat                = errors.New("invalid field format")                                  // 2003
	ErrInvalidStatusTransition           = errors.New("invalid status transition")                             // 2004
	ErrInvalidQueryRequestParam          = errors.New("invalid query request param")                           // 2005
	ErrInvalidRequestParam               = errors.New("invalid request param")                                 // 2006
	ErrHashPassword                      = errors.New("failed to hash password")                               // 2007
	ErrConvertFileHeaderToBase64         = errors.New("failed to convert file header to base64")               // 2008
	ErrOpenFileContext                   = errors.New("failed to open file context")                           // 2009
	ErrUserIdDoesNotMatchWithYourAccount = errors.New("user id of this task does not match with your account") // 2010
	ErrOtherError                        = errors.New("other error")                                           // 2012

	// Task Resource
	ErrTaskNotFound      = errors.New("task not found")      // 3001
	ErrTaskAlreadyExists = errors.New("task already exists") // 3002

	// User Resource
	ErrUserNotFound      = errors.New("user not found")        // 4001
	ErrPasswordIncorrect = errors.New("password is incorrect") // 4002
	ErrUserAlreadyExists = errors.New("user already exists")   // 4003

	// Internal
	ErrInternalServerError       = errors.New("internal server error")                   // 5001
	ErrServiceUnavailable        = errors.New("service unavailable")                     // 5002
	ErrGetCurrentTimeWithRFC3339 = errors.New("failed to get current time with RFC3339") // 5003
)

var ErrorMapWithCode = map[error]ErrorType{
	// Token-related errors
	ErrExpiredToken:                      CodeTokenExpired,                      // 1001
	ErrInvalidToken:                      CodeTokenInvalid,                      // 1002
	ErrFailedToVerifyToken:               CodeTokenVerificationFailed,           // 1003
	ErrAuthHeaderMissing:                 CodeAuthHeaderMissing,                 // 1004
	ErrAuthHeaderFormatInvalid:           CodeAuthHeaderFormatInvalid,           // 1005
	ErrAuthHeaderMissingBearer:           CodeAuthHeaderMissingBearer,           // 1006
	ErrUserIdDoesNotMatchWithYourAccount: CodeUserIdDoesNotMatchWithYourAccount, // 1007
	ErrUnauthorized:                      CodeUnauthorized,                      // 1008

	// Input & validation
	ErrInvalidRequestBody:                CodeInvalidRequestBody,                // 2001
	ErrMissingRequiredFields:             CodeMissingRequiredFields,             // 2002
	ErrInvalidFieldFormat:                CodeInvalidFieldFormat,                // 2003
	ErrInvalidStatusTransition:           CodeInvalidStatusTransition,           // 2004
	ErrInvalidQueryRequestParam:          CodeInvalidQueryRequestParam,          // 2005
	ErrInvalidRequestParam:               CodeInvalidRequestParam,               // 2006
	ErrHashPassword:                      CodeHashPassword,                      // 2007
	ErrConvertFileHeaderToBase64:         CodeConvertFileHeaderToBase64,         // 2008
	ErrOpenFileContext:                   CodeOpenFileContext,                   // 2009
	ErrUserIdDoesNotMatchWithYourAccount: CodeUserIdDoesNotMatchWithYourAccount, // 2010
	ErrOtherError:                        CodeOtherError,                        // 2012

	// Task Resource
	ErrTaskNotFound:      CodeTaskNotFound,      // 3001
	ErrTaskAlreadyExists: CodeTaskAlreadyExists, // 3002

	// User Resource
	ErrUserNotFound:      CodeUserNotFound,      // 4001
	ErrPasswordIncorrect: CodePasswordIncorrect, // 4002
	ErrUserAlreadyExists: CodeUserAlreadyExists, // 4003

	// Internal
	ErrInternalServerError:       CodeInternalServerError,       // 5001
	ErrServiceUnavailable:        CodeServiceUnavailable,        // 5002
	ErrGetCurrentTimeWithRFC3339: CodeGetCurrentTimeWithRFC3339, // 5003
}

var ErrorMapWithStatusCode = map[error]int{
	// Token-related errors
	ErrExpiredToken:                      http.StatusUnauthorized, // 1001
	ErrInvalidToken:                      http.StatusUnauthorized, // 1002
	ErrFailedToVerifyToken:               http.StatusUnauthorized, // 1003
	ErrAuthHeaderMissing:                 http.StatusUnauthorized, // 1004
	ErrAuthHeaderFormatInvalid:           http.StatusUnauthorized, // 1005
	ErrAuthHeaderMissingBearer:           http.StatusUnauthorized, // 1006
	ErrUserIdDoesNotMatchWithYourAccount: http.StatusUnauthorized, // 1007
	ErrUnauthorized:                      http.StatusUnauthorized, // 1008

	// Input & validation
	ErrInvalidRequestBody:                http.StatusBadRequest,   // 2001
	ErrMissingRequiredFields:             http.StatusBadRequest,   // 2002
	ErrInvalidFieldFormat:                http.StatusBadRequest,   // 2003
	ErrInvalidStatusTransition:           http.StatusBadRequest,   // 2004
	ErrInvalidQueryRequestParam:          http.StatusBadRequest,   // 2005
	ErrInvalidRequestParam:               http.StatusBadRequest,   // 2006
	ErrHashPassword:                      http.StatusBadRequest,   // 2007
	ErrConvertFileHeaderToBase64:         http.StatusBadRequest,   // 2008
	ErrOpenFileContext:                   http.StatusBadRequest,   // 2009
	ErrUserIdDoesNotMatchWithYourAccount: http.StatusUnauthorized, // 2010
	ErrOtherError:                        http.StatusBadRequest,   // 2012

	// Task Resource
	ErrTaskNotFound:      http.StatusNotFound, // 3001
	ErrTaskAlreadyExists: http.StatusConflict, // 3002

	// User Resource
	ErrUserNotFound:      http.StatusNotFound,     // 4001
	ErrPasswordIncorrect: http.StatusUnauthorized, // 4002
	ErrUserAlreadyExists: http.StatusConflict,     // 4003

	// Internal
	ErrInternalServerError:       http.StatusInternalServerError, // 5001
	ErrServiceUnavailable:        http.StatusServiceUnavailable,  // 5002
	ErrGetCurrentTimeWithRFC3339: http.StatusInternalServerError, // 5003
}
