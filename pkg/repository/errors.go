package repository

import "errors"

var (
	ErrDoesNotExist       = errors.New("does not exist")
	ErrInvalidBanDuration = errors.New("invalid ban duration")
)
