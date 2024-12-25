package repository

import (
	"errors"
	"fmt"

	"github.com/ShatteredRealms/chat-service/pkg/common"
)

var (
	ErrDoesNotExist       = errors.New("does not exist")
	ErrInvalidBanDuration = fmt.Errorf("%w: invalid ban duration", common.ErrRequestInvalid)
	ErrNoUpdates          = errors.New("no updates")
	ErrEmptyMessage       = fmt.Errorf("%w: empty message", common.ErrRequestInvalid)

	ErrInvalidTimeRange = fmt.Errorf("%w: invalid time range", common.ErrRequestInvalid)
	ErrInvalidLimit     = fmt.Errorf("%w: invalid limit", common.ErrRequestInvalid)
	ErrInvalidOffset    = fmt.Errorf("%w: invalid offset", common.ErrRequestInvalid)

	ErrNilId      = fmt.Errorf("%w: id is nil", common.ErrRequestInvalid)
	ErrNonEmptyId = fmt.Errorf("%w: id is not empty", common.ErrRequestInvalid)
)
