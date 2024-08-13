package contracts

import (
	"errors"
	"strings"
)

const (
	POSIX_MAX_HOST_LEN = 255

	ALLOWED_EMAIL_NAME_CHARSET = "0123456789qwertyuiopasdfghjklzxcvbnm+-_~" // Does not include upper case.
	ALLOWED_HOSTNAME_CHARSET   = "0123456789qwertyuiopasdfghjklzxcvbnm.-"   // Does not include upper case.
)

var (
	ErrInvalidName    = errors.New("invalid name")
	ErrInvalidSurname = errors.New("invalid surname")
	ErrInvalidEmail   = errors.New("invalid email address")
)

type RegisterReqCtx struct {
	Name    string
	Surname string
	Email   string
}

func (ctx *RegisterReqCtx) Validate() (err error) {
	if len(ctx.Name) < 2 || len(ctx.Name) > 20 {
		return ErrInvalidName
	}

	if len(ctx.Surname) < 2 || len(ctx.Surname) > 40 {
		return ErrInvalidSurname
	}

	segments := strings.Split(ctx.Email, "@")
	if len(segments) != 2 {
		return ErrInvalidEmail
	}

	if len(segments[1]) > POSIX_MAX_HOST_LEN {
		return ErrInvalidEmail
	}

	for _, ch := range strings.ToLower(segments[0]) {
		if !strings.Contains(ALLOWED_EMAIL_NAME_CHARSET, string(ch)) {
			return ErrInvalidEmail
		}
	}

	for _, ch := range strings.ToLower(segments[1]) {
		if !strings.Contains(ALLOWED_HOSTNAME_CHARSET, string(ch)) {
			return ErrInvalidEmail
		}
	}

	return nil
}
