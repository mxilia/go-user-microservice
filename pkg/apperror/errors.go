package apperror

import "errors"

var (
	// ------------------------
	// 400 Bad Request
	// ------------------------
	ErrInvalidInput   = errors.New("invalid input")
	ErrInvalidState   = errors.New("invalid state")
	ErrValidation     = errors.New("validation failed")
	ErrInvalidID      = errors.New("invalid id")
	ErrInvalidFormat  = errors.New("invalid format")
	ErrRequiredField  = errors.New("required field")
	ErrOutOfRange     = errors.New("value out of range")
	ErrMalformedBody  = errors.New("malformed request body")
	ErrInvalidRequest = errors.New("invalid request")

	// ------------------------
	// 401 Unauthorized
	// ------------------------
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidToken   = errors.New("invalid token")
	ErrExpiredToken   = errors.New("expired token")
	ErrInvalidLogin   = errors.New("invalid credentials")
	ErrSessionExpired = errors.New("session expired")
	ErrAuthentication = errors.New("authentication failed")

	// ------------------------
	// 403 Forbidden
	// ------------------------
	ErrForbidden           = errors.New("forbidden")
	ErrOperationNotAllowed = errors.New("operation not allowed")
	ErrPermissionDenied    = errors.New("permission denied")

	// ------------------------
	// 404 Not Found
	// ------------------------
	ErrNotFound = errors.New("resource not found")

	// ------------------------
	// 405 Method Not Allowed
	// ------------------------
	ErrMethodNotAllowed = errors.New("method not allowed")

	// ------------------------
	// 409 Conflict
	// ------------------------
	ErrAlreadyExists = errors.New("resource already exists")
	ErrConflict      = errors.New("resource conflict")
	ErrDuplicate     = errors.New("duplicate resource")

	// ------------------------
	// 410 Gone
	// ------------------------
	ErrArchived = errors.New("resource archived")
	ErrDeleted  = errors.New("resource deleted")

	// ------------------------
	// 422 Unprocessable Entity
	// ------------------------
	ErrBusinessRuleViolation = errors.New("business rule violation")
	ErrInvalidTransition     = errors.New("invalid state transition")

	// ------------------------
	// 429 Too Many Requests
	// ------------------------
	ErrLimitExceeded = errors.New("limit exceeded")
	ErrRateLimited   = errors.New("rate limited")

	// ------------------------
	// 500 Internal Server Error
	// ------------------------
	ErrInternal      = errors.New("internal error")
	ErrUnexpected    = errors.New("unexpected error")
	ErrConfiguration = errors.New("configuration error")
	ErrTransaction   = errors.New("transaction failed")

	// ------------------------
	// 501 Not Implemented
	// ------------------------
	ErrNotImplemented = errors.New("not implemented")

	// ------------------------
	// 502 Bad Gateway
	// ------------------------
	ErrExternalService = errors.New("external service error")
	ErrDependency      = errors.New("dependency failure")

	// ------------------------
	// 503 Service Unavailable
	// ------------------------
	ErrServiceUnavailable = errors.New("service unavailable")

	// ------------------------
	// 504 Gateway Timeout
	// ------------------------
	ErrTimeout = errors.New("operation timeout")
)
