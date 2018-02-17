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

// Function to insert elements in a circular queue
func (q *Queue) InsertCircular(element int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if (q.rear+1)%q.capacity == q.front {
		return errors.New("Queue Capacity Full")
	} else {
		if q.front == -1 {
			q.front = 0
		}
		q.rear = (q.rear+1)%q.capacity
		q.arr[q.rear] = element
		return nil
	}
}

// Function to delete elements from the circular queue
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

// Function to insert into a priority queue
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