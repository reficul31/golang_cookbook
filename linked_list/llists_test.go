package llists

import "testing"

// Test the new node memory allocation function of NewLinkedListNode
func TestNewLinkedListNode(testCase *testing.T) {
	testCase.Log("To test that the new linked list node is allocated the correct value")

	new := NewLinkedListNode(1)
	if new.data != 1 || new.next != nil {
		testCase.Errorf("LList Error: Node not assigned correct value")
	} 
}

// Test the InsertAtHead function of the Linked List
func TestInsertAtHead(testCase *testing.T) {
	testCase.Log("To test that the new linked list node is inserted at the head of the linked list")
	head := NewLinkedListNode(-1)

	head.InsertAtHead(1)
	if head.next.data != 1 || head.next.next != nil {
		testCase.Errorf("LList Error: Node not assigned correct value")
	}

	head.InsertAtHead(2)
	if head.next.data != 2 || head.next.next == nil {
		testCase.Errorf("LList Error: Node not assigned correct value")
	}	
}

// Test the InsertAtBack function of the Linked List
func TestInsertAtBack(testCase *testing.T) {
	testCase.Log("To test that the new linked list node is inserted at the back of the linked list")
	head := NewLinkedListNode(-1)

	head.InsertAtBack(1)
	if head.next.data != 1 || head.next.next != nil {
		testCase.Errorf("LList Error: Node not assigned correct value")
	}

	head.InsertAtBack(2)
	if head.next.next.data != 2 || head.next.next.next != nil {
		testCase.Errorf("LList Error: Node not assigned correct value")
	}	
}

// Test the DeleteFromHead function of the Linked List
func TestDeleteFromHead(testCase *testing.T) {
	testCase.Log("To test the linked list node is deleted from the head of the linked list")
	head := NewLinkedListNode(-1)

	head.InsertAtHead(1)
	head.InsertAtHead(2)
	head.InsertAtHead(3)

	element, err := head.DeleteFromHead()
	if element != 3 || err != nil {
		testCase.Errorf("LList Error: Wrong Node was deleted")
	}

	_, _ = head.DeleteFromHead()
	_, _ = head.DeleteFromHead()

	if _, err = head.DeleteFromHead() ; err == nil {
		testCase.Errorf("LList Error: Linked List underflow not reported")
	}
}

// Test the DeleteFromBack function of the Linked List
func TestDeleteFromBack(testCase *testing.T) {
	testCase.Log("To test the linked list node is deleted from the back of the linked list")
	head := NewLinkedListNode(-1)

	head.InsertAtBack(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)

	element, err := head.DeleteFromBack()
	if element != 3 || err != nil {
		testCase.Errorf("LList Error: Wrong Node was deleted")
	}

	_, _ = head.DeleteFromBack()
	_, _ = head.DeleteFromBack()

	if _, err = head.DeleteFromBack() ; err == nil {
		testCase.Errorf("LList Error: Linked List underflow not reported")
	}
}