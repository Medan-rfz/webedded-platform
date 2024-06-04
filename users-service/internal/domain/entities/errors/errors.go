package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")

	ErrParseToken          = errors.New("error token parse")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)
