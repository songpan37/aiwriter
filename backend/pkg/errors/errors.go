package errors

import "errors"

var (
	ErrNotFound         = errors.New("resource not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrUserExists       = errors.New("user already exists")
	ErrEmailExists      = errors.New("email already exists")
	ErrTokenExpired     = errors.New("token expired")
	ErrTokenInvalid     = errors.New("token invalid")
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrAIServiceFailed  = errors.New("AI service failed")
	ErrPublishFailed    = errors.New("publish failed")
)
