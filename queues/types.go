package queues

import "sync"

type Queue struct {
	capacity int
	front    int
	rear     int
	lock     sync.Mutex
	arr      []int
}