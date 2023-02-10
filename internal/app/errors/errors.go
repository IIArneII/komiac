package errors

import "errors"

var (
	ErrNotFound = errors.New("not found")
	BadUUID     = errors.New("bad uuid")
)
