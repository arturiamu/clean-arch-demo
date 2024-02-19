package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrWrongPassword      = errors.New("wrong password")
	ErrInvalidAccessToken = errors.New("invalid access token")
)
