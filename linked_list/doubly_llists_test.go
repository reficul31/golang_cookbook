package llists

import "testing"

// Test the new node memory allocation function of NewDoublyLinkedListNode
func TestNewDoublyLinkedListNode(testCase *testing.T) {
	testCase.Log("To test the new doubly linked list node is allocated the correct value")

	new := NewDoublyLinkedListNode(1)
	if new.data != 1 || new.next != nil || new.prev != nil {
		testCase.Errorf("Doubly LList Error: Node not assigned correct value")
	}
}

// Test the InsertAtHeadDLL function of the Doubly Linked List
func TestInsertAtHeadDLL(testCase *testing.T) {
	testCase.Log("To test that the new doubly linked list node is inserted at the head of the linked list")
	head := NewDoublyLinkedListNode(-1)

	head.InsertAtHeadDLL(1)
	if head.next.data != 1 || head.next.next != nil || head.next.prev != head{
		testCase.Errorf("Doubly LList Error: Node not assigned the correct value")
	}

	head.InsertAtHeadDLL(2)
	if head.next.data != 2 || head.next.next == nil || head.next.prev != head {
		testCase.Errorf("Doubly LList Error: Node not assigned the correct value")
	}

	node2 := head.next.next
	if node2.data != 1 || node2.prev != head.next {
		testCase.Errorf("Doubly LList Error: Node not assigned the correct value")
	}
}

// Test the InsertAtBackDLL function of the Doubly Linked List
func TestInsertAtBackDLL(testCase *testing.T) {
	testCase.Log("To test that the new doubly linked list node is inserted at the back of the linked list")
	head := NewDoublyLinkedListNode(-1)

	head.InsertAtBackDLL(1)
	if head.next.data != 1 || head.next.next != nil || head.next.prev != head {
		testCase.Errorf("Doubly LList Error: Node not assignmed the correct value")
	}

	head.InsertAtBackDLL(2)
	node2 := head.next.next
	if node2.data != 2 || node2.prev != head.next || node2.next != nil || head.next.next != node2 {
		testCase.Errorf("Doubly LList Error: Node not assigned the correct value")
	}
}

// Test the DeleteFromHeadDLL function of the Doubly Linked List
func TestDeleteFromHeadDLL(testCase *testing.T) {
	testCase.Log("To test the doubly linked list node is deleted from the head of the linked list")
	head := NewDoublyLinkedListNode(-1)

	head.InsertAtHeadDLL(1)
	head.InsertAtHeadDLL(2)
	head.InsertAtHeadDLL(3)

	element, err := head.DeleteFromHeadDLL()
	if element != 3 || err != nil {
		testCase.Errorf("Doubly LList Error: Wrong Node was deleted")
	}
	if head.next.data != 2 || head.next.prev != head {
		testCase.Errorf("Doubly LList Error: Delete operation wasn't successful")
	}

	_, _ = head.DeleteFromHeadDLL()
	_, _ = head.DeleteFromHeadDLL()

	if _, err = head.DeleteFromHeadDLL() ; err == nil {
		testCase.Errorf("Doubly LList Error: Linked List underflow not reported")
	}
}

// Test the DeleteFromBackDLL function of the Doubly Linked List
func TestDeleteFromBackDLL(testCase *testing.T) {
	testCase.Log("To test the doubly linked list noe is deleted from the back of the linked list")
	head := NewDoublyLinkedListNode(-1)

	head.InsertAtBackDLL(1)
	head.InsertAtBackDLL(2)
	head.InsertAtBackDLL(3)

	element, err := head.DeleteFromBackDLL()
	if element != 3 || err != nil {
		testCase.Errorf("Doubly LList Error: Wrong Node was deleted")
	}
	node2 := head.next.next
	if node2.next != nil || node2.data != 2 {
		testCase.Errorf("Doubly LList Error: Delete operation wasn't successful")
	}

	_, _ = head.DeleteFromBackDLL()
	_, _ = head.DeleteFromBackDLL()

	if _, err = head.DeleteFromBackDLL() ; err == nil {
		testCase.Errorf("Doubly LList Error: Linked List underflow not reported")
	}
}