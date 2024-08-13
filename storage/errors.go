package storage

import "errors"

var (
	ErrUserInvalid  = errors.New("invalid user")
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")

	ErrSpinInvalid  = errors.New("invalid spin")
	ErrSpinNotFound = errors.New("spin not found")
)
