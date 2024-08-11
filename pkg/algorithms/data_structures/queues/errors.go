package queues

import "errors"

var (
	ErrQueueEmpty   = errors.New("queue is empty")
	ErrInvalidPos   = errors.New("invalid position or list is empty")
	ErrElemNotFound = errors.New("element not found in queue")
)
