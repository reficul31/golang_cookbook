package queues

import "testing"

var err error

// Test the queue capacity allocation function NewQueue
func TestQueueCapacityAllocation(testCase *testing.T) {
	testCase.Log("To test that the queue size is allocated correctly")

	if q, err := NewQueue(20); err == nil {
		if len(q.arr) != 20 {
			testCase.Errorf("Queue Error: Queue size not equal to capacity")
		}
	} else {
		testCase.Errorf("Queue Error: %s", err)
	}
}

// Test the Insert function of the Queue
func TestQueueInsert(testCase *testing.T) {
	testCase.Log("To test that the queue insert function inserts numbers in the queue")
	
	q, _ := NewQueue(2)
	if err = q.Insert(1) ; err != nil {
		testCase.Errorf("Queue Error: %s", err)
	} else {
		if q.arr[q.rear] != 1 {
			testCase.Errorf("Queue Error: Element not inserted")
		}
	}

	err = q.Insert(2)
	if err = q.Insert(3) ; err == nil {
		testCase.Errorf("Queue Error: Overflow not reported")
	}
}

// Test the Delete function of the Queue
func TestQueueDelete(testCase *testing.T) {
	testCase.Log("To test that the queue delete function deletes the number from the queue")
	var element int

	q, _ := NewQueue(3)
	_ = q.Insert(3)
	_ = q.Insert(2)
	_ = q.Insert(1)

	if element, err = q.Delete() ; err != nil {
		testCase.Errorf("Queue Error: %s", err)
	}
	if element != 3 {
		testCase.Errorf("Queue Error: Element not deleted")
	}
	
	_, err = q.Delete()
	_, err = q.Delete()

	if _, err = q.Delete() ; err == nil {
		testCase.Errorf("Queue Error: Queue Underflow not reported")
	}
}