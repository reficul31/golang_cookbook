package queues

// InsertCircular inserts an element in a circular queue
func (q *Queue) InsertCircular(element int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if (q.rear+1)%q.capacity == q.front {
		return ErrQueueFull
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
		return -1, ErrQueueEmpty
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