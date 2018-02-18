package stacks

import "errors"

var (
	// ErrStackCapacity is returned when stack capacity is not a positive integer
	ErrStackCapacity = errors.New("Stack Capacity must be a positive integer")

	// ErrStackFull is returned when an insertion if performed on a full stack
	ErrStackFull     = errors.New("Stack Full")

	// ErrStackEmpty is returned when a deletion is performed on an empty stack
	ErrStackEmpty    = errors.New("Stack Empty")
)