package stacks

import "sync"

// Stack is the structure defining the stack
type Stack struct {
	capacity int
	top      int
	lock     sync.Mutex
	arr      []int
}