package queues

// InsertPriority inserts an element based on the priority function provided to it
func (q *Queue) InsertPriority(element int, priorityDecider func(int, int) bool) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.rear == q.capacity-1 {
		return ErrQueueFull
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