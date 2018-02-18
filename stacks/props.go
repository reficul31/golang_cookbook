package stacks

import (
	"sync"
	"errors"
)

// NewStack is used to return a pointer to a newly allocated stack
func NewStack(capacity int) (*Stack, error) {
	if capacity >= 0 {
		return &Stack{
			capacity,
			-1,
			sync.Mutex{}, 
			make([]int, capacity),
		}, nil
	}

	return nil, errors.New("Stack Capacity cannot be negative")
}

// Push is used to push elements on top of stack
func (s *Stack) Push(element int) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top != s.capacity-1 {
		s.top = s.top + 1
		s.arr[s.top] = element
		return nil
	}
	
	return errors.New("Stack Capacity Full")
}

// Pop is used to pop elements from top of the stack
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