package queues

import "errors"

var (
	ErrQueueEmpty        = errors.New("queue is empty")
	ErrInvalidPosQueue   = errors.New("invalid position or queue is empty")
	ErrElemNotFoundQueue = errors.New("element not found in queue")

	ErrDequeEmpty        = errors.New("deque is empty")
	ErrInvalidPosDeque   = errors.New("invalid position or deque is empty")
	ErrElemNotFoundDeque = errors.New("element not found in deque")
)
