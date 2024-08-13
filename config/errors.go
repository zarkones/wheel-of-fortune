package config

import "errors"

var (
	ErrInvalidMinPrize = errors.New("invalid minimum prize score")
	ErrInvalidMaxPrize = errors.New("invalid maximum prize score")
)
