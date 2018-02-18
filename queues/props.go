package queues

import (
	"sync"
	"errors"
)

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

	return nil, errors.New("Queue Capacity cannot be negative")
}

// Insert is used to insert an element in the queue
func (q *Queue) Insert(element int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.rear == q.capacity-1 {
		return errors.New("Queue Capacity Full")
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

// InsertCircular inserts an element in a circular queue
func (q *Queue) InsertCircular(element int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if (q.rear+1)%q.capacity == q.front {
		return errors.New("Queue Capacity Full")
	}
	
	if q.front == -1 {
		q.front = 0
	}
	q.rear = (q.rear+1)%q.capacity
	q.arr[q.rear] = element
	return nil
}

// DeleteCircular deletes an element from a circular queue
func (q *Queue) DeleteCircular() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	var element int

	if q.rear == -1 && 	q.front == -1 {
		return -1, errors.New("Queue Empty")
	} else if q.front == q.rear {
			element = q.arr[q.front]
			q.front = -1
			q.rear = -1
	} else {
		element = q.arr[q.front]
		q.front = (q.front+1)%q.capacity
	}
	return element, nil
}

// InsertPriority inserts an element based on the priority function provided to it
func (q *Queue) InsertPriority(element int, priorityDecider func(int, int) bool) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.rear == q.capacity-1 {
		return errors.New("Queue Capacity Full")
	} else if q.rear == -1 && q.front == -1 {
		q.front = q.front+1
		q.rear = q.rear+1
		q.arr[q.rear] = element
		return nil
	} else {
		var i int
		for i=q.rear+1 ; i>=q.front && !priorityDecider(q.arr[i-1], element); i-- {
			q.arr[i] = q.arr[i-1]
		}
		q.arr[i] = element
		q.rear = q.rear+1
		return nil
	}
}

// InsertFront inserts element in the front of the queue
func (q *Queue) InsertFront(element int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.front == -1 {
		q.front = 0
		q.rear = 0
		q.arr[q.front] = element
		return nil
	} else if q.front == 0 {
		return errors.New("Cannot insert from front")
	} else {
		q.front = q.front - 1
		q.arr[q.front] = element
		return nil
	}
}

// DeleteRear deletes from the rear of the queue
func (q *Queue) DeleteRear() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	var element int

	if q.front == -1 && q.rear == -1 {
		return -1, errors.New("Queue Empty")
	} else if q.front == q.rear {
		element = q.arr[q.rear]
		q.front = -1
		q.rear = -1
		return element, nil
	} else {
		element = q.arr[q.rear]
		q.rear = q.rear - 1
		return element, nil
	}
}