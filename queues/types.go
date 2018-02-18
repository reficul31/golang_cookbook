package queues

import "sync"

// Queue is the structure defining the queue
type Queue struct {
	capacity int
	front    int
	rear     int
	lock     sync.Mutex
	arr      []int
}