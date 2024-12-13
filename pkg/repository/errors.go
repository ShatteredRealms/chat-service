package repository

import "errors"

var (
	ErrDoesNotExist       = errors.New("does not exist")
	ErrInvalidBanDuration = errors.New("invalid ban duration")
	ErrNoUpdates          = errors.New("no updates")
	ErrEmptyMessage       = errors.New("empty message")

	ErrInvalidTimeRange = errors.New("invalid time range")
	ErrInvalidLimit     = errors.New("invalid limit")
	ErrInvalidOffset    = errors.New("invalid offset")
)
