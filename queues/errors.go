package queues

import "errors"

var (
	// ErrQueueCapacity is returned when queue capacity is not a positive integer
	ErrQueueCapacity = errors.New("Queue Capacity must be a positive integer")

	// ErrQueueEmpty is returned when a deletion operation is performed on an empty queue
	ErrQueueEmpty    = errors.New("Queue Empty")

	// ErrQueueFull is returned when an insertion operation is performed on a full queue
	ErrQueueFull     = errors.New("Queue Full")

	// ErrInsertFront is returned when an InsertFront operation is performed illegally
	ErrInsertFront   = errors.New("Queue cannot insert at the front")
)