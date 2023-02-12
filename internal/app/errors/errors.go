package errors

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrBadUUID     = errors.New("bad uuid")
	ErrEntityIsNil = errors.New("entity is nil")
)
