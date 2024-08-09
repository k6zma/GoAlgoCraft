package linked_lists

import "errors"

var (
	ErrInvalidPos    = errors.New("invalid position or list is empty")
	ErrPosOutOfRange = errors.New("position out of range")
	ErrListEmpty     = errors.New("list is empty")
	ErrValueNotFound = errors.New("value not found")
)
