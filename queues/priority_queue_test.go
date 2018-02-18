package queues

import "testing"

// Test the priority insert of the queue
func TestPriorityInsert(testCase *testing.T) {
	testCase.Log("To test that the priority queue insert function inserts numbers in the queue")

	less := func(a int, b int) bool {
		if a>b {
			return false
		}

		return true
	}
	
	q, _ := NewQueue(5)
	if err = q.InsertPriority(1, less) ; err != nil {
		testCase.Errorf("Queue Error: %s", err)
	} else {
		if q.arr[q.rear] != 1 {
			testCase.Errorf("Queue Error: Element not inserted")
		}
	}

	_ = q.InsertPriority(5, less)
	if q.arr[q.rear] != 5 {
		testCase.Errorf("Queue Error: Priority Function not applied")
	}
	_ = q.InsertPriority(3, less)
	if q.arr[q.rear] != 5 {
		testCase.Errorf("Queue Error: Priority Function not applied")
	}
	_ = q.InsertPriority(6, less)
	if q.arr[q.rear] != 6 {
		testCase.Errorf("Queue Error: Priority Function not applied")
	}
	_ = q.InsertPriority(2, less)
	if q.arr[q.rear] != 6 {
		testCase.Errorf("Queue Error: Priority Function not applied")
	}
	if err = q.InsertPriority(3, less) ; err == nil {
		testCase.Errorf("Queue Error: Overflow not reported")
	}
}
