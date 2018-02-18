package queues

import "testing"

// Test the Insert Front function of the Queue
func TestQueueInsertFront(testCase *testing.T) {
	testCase.Log("To test that the queue insert front function inserts numbers in the front of the queue")
	
	q, _ := NewQueue(2)
	if err = q.InsertFront(1) ; err != nil {
		testCase.Errorf("Queue Error: %s", err)
	} else {
		if q.arr[q.rear] != 1 {
			testCase.Errorf("Queue Error: Element not inserted")
		}
	}

	if err = q.InsertFront(3) ; err == nil {
		testCase.Errorf("Queue Error: Front insertion error not reported")
	}

	_, _ = q.Delete()
	if err = q.InsertFront(3) ; err != nil {
		testCase.Errorf("Queue Error: %s", err)
	} else {
		if q.arr[q.rear] != 3 {
			testCase.Errorf("Queue Error: Element not inserted")
		}
	}
}

// Test the Delete Rear function of the Queue
func TestQueueDeleteRear(testCase *testing.T) {
	testCase.Log("To test that the queue delete rear function deletes the number from the rear of the queue")
	var element int

	q, _ := NewQueue(3)
	_ = q.Insert(3)
	_ = q.Insert(2)
	_ = q.Insert(1)

	if element, err = q.DeleteRear() ; err != nil {
		testCase.Errorf("Queue Error: %s", err)
	}
	if element != 1 {
		testCase.Errorf("Queue Error: Element not deleted")
	}
	
	_, err = q.DeleteRear()
	_, err = q.DeleteRear()

	if _, err = q.DeleteRear() ; err == nil {
		testCase.Errorf("Queue Error: Queue Underflow not reported")
	}
}