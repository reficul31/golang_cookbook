package stacks

import (
	"sync"
	"errors"
)

// Function to create a new Stack
func NewStack(capacity int) (*Stack, error) {
	if capacity >= 0 {
		return &Stack{
			capacity,
			-1,
			sync.Mutex{}, 
			make([]int, capacity),
		}, nil
	} else {
		return nil, errors.New("Stack Capacity cannot be negative")
	}
}

// Function to push elements into the stack
func (s *Stack) Push(element int) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top != s.capacity-1 {
		s.top = s.top + 1
		s.arr[s.top] = element
		return nil
	} else {
		return errors.New("Stack Capacity Full")
	}
}

// Function to pop elements from the stack
func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == -1 {
		return -1, errors.New("Stack Empty")
	}

	element := s.arr[s.top]
	s.top = s.top - 1
	return element, nil
}