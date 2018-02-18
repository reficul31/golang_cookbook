package llists

import "errors"

var (
	// ErrLinkedListEmpty is returned when a delete function is performed on an empty linked list
	ErrLinkedListEmpty = errors.New("Linked List Empty")
)