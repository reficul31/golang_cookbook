package queues

import (
	"sync"
	"errors"
)

// Function to create a new Queue
func NewQueue(capacity int) (*Queue, error) {
	if capacity >= 0 {
		return &Queue{
			capacity,
			-1,
			-1,
			sync.Mutex{},
			make([]int, capacity),
		}, nil
	} else {
		return nil, errors.New("Queue Capacity cannot be negative")
	}
}

// Function to insert elements into the queue
func (q *Queue) Insert(element int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.rear == q.capacity-1 {
		return errors.New("Queue Capacity Full")
	} else {
		if q.front == -1 {
			q.front = 0
		}
		q.rear = q.rear+1
		q.arr[q.rear] = element
		return nil
	}
}

// Function to delete elements from the queue
func (q *Queue) Delete() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	var element int

	if q.rear == -1 && q.front == -1 {
		return -1, errors.New("Queue Empty")
	} else if q.front == q.rear {
		element = q.arr[q.front]
		q.front = -1
		q.rear = -1
	} else {
		element = q.arr[q.front]
		q.front = q.front + 1
	}
	return element, nil
}