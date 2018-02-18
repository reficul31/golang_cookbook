package queues

import "sync"

// NewQueue is used to return a pointer to a newly allocated queue
func NewQueue(capacity int) (*Queue, error) {
	if capacity >= 0 {
		return &Queue{
			capacity,
			-1,
			-1,
			sync.Mutex{},
			make([]int, capacity),
		}, nil
	}

	return nil, ErrQueueCapacity
}

// Insert is used to insert an element in the queue
func (q *Queue) Insert(element int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.rear == q.capacity-1 {
		return ErrQueueFull
	}

	if q.front == -1 {
		q.front = 0
	}
	q.rear = q.rear+1
	q.arr[q.rear] = element
	return nil
}

// Delete is used to delete an element from the queue
func (q *Queue) Delete() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	var element int

	if q.rear == -1 && q.front == -1 {
		return -1, ErrQueueEmpty
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