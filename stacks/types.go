package stacks

import "sync"

type Stack struct {
	capacity int
	top      int
	lock     sync.Mutex
	arr      []int
}