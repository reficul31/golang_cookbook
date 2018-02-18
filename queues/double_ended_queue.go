package queues

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
		return ErrInsertFront
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
		return -1, ErrQueueEmpty
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