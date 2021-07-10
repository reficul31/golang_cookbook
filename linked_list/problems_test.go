package llists

import (
	"testing"
)

// Test the delete duplicate using Hash Maps of Linked List
func TestDeleteDuplicateUsingHashMaps(testCase *testing.T) {
	testCase.Log("To test the delete duplicate using Hash Maps in Linked List")

	head := NewLinkedListNode(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)
	head.InsertAtBack(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)

	new := DeleteDuplicateUsingHashMaps(head)
	if new.data != 1 || new.next.data != 2 || new.next.next.data != 3 || new.next.next.next.next != nil {
		testCase.Errorf("Linked List Error: Linked list contains duplicate elements")
	}
}

// Test the delete duplicate using iteration of Linked List
func TestDeleteDuplicateUsingIteration(testCase *testing.T) {
	testCase.Log("To test the delete duplicate using iteration in Linked List")

	head := NewLinkedListNode(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)
	head.InsertAtBack(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)

	new := DeleteDuplicateUsingIteration(head)
	if new.data != 1 || new.next.data != 2 || new.next.next.data != 3 || new.next.next.next.next != nil {
		testCase.Errorf("Linked List Error: Linked list contains duplicate elements")
	}
}

// Test the nth to the last node in a Linked List
func TestNthToLastNode(testCase *testing.T) {
	testCase.Log("To test the nth to the last node in a Linked List")

	head := NewLinkedListNode(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)
	head.InsertAtBack(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)

	nthNode := NthToLastNode(head, 2)
	if nthNode.data != 2 {
		testCase.Errorf("Linked List Error: Linked list returned the wrong nth last node")
	}
}

// Test the partition at N in a Linked List
func TestPartitionAtN(testCase *testing.T) {
	testCase.Log("To test the partition at N in a Linked List")

	head := NewLinkedListNode(4)
	head.InsertAtBack(6)
	head.InsertAtBack(2)
	head.InsertAtBack(1)
	head.InsertAtBack(3)
	head.InsertAtBack(7)

	new := PartitionAtN(head, 5)
	if new.data != 3 || new.next.data != 1 || new.next.next.data != 2 || new.next.next.next.data != 4 || new.next.next.next.next.data != 6 || new.next.next.next.next.next.data != 7 || new.next.next.next.next.next.next != nil {
		testCase.Errorf("Linked List Error: Linked list partitioned along the wrong node")
	}
}

// Test the addition of two numbers using linked list
func TestAddTwoNumbersUsingLinkedList(testCase *testing.T) {
	testCase.Log("To test the addition of two numbers using Linked List")

	additive1 := NewLinkedListNode(1)
	additive1.InsertAtBack(2)
	additive1.InsertAtBack(3)
	additive1.InsertAtBack(4)

	additive2 := NewLinkedListNode(1)
	additive2.InsertAtBack(2)
	additive2.InsertAtBack(3)
	additive2.InsertAtBack(4)

	sum := AddTwoNumbersUsingLinkedList(additive1, additive2)
	if sum.data != 2 || sum.next.data != 4 || sum.next.next.data != 6 || sum.next.next.next.data != 8 || sum.next.next.next.next != nil {
		testCase.Errorf("Linked List Error: Linked list addition operation error")
	}
}

// Test whether the linked list is a palindrome
func TestIsPalindrome(testCase *testing.T) {
	testCase.Log("To test whether the linked list is a palindrome")

	palindrome := NewLinkedListNode(1)
	palindrome.InsertAtBack(2)
	palindrome.InsertAtBack(3)
	palindrome.InsertAtBack(2)
	palindrome.InsertAtBack(1)

	if !IsPalindrome(palindrome) {
		testCase.Errorf("Linked List Error: Linked list is a palindrome but function returned false")
	}

	notPalindrome := NewLinkedListNode(1)
	notPalindrome.InsertAtBack(2)
	notPalindrome.InsertAtBack(3)
	notPalindrome.InsertAtBack(2)
	notPalindrome.InsertAtBack(4)

	if IsPalindrome(notPalindrome) {
		testCase.Errorf("Linked List Error: Linked list is not a palindrome but function returned true")
	}
}

// Test to find the intersection node in a linked list
func TestFindIntersectionNode(testCase *testing.T) {
	testCase.Log("To test the intersection node in a linked list")

	list1 := NewLinkedListNode(1)
	list1.InsertAtBack(2)
	list1.InsertAtBack(3)

	list2 := NewLinkedListNode(1)
	list2.InsertAtBack(2)
	list2.InsertAtBack(3)

	if FindIntersectionNode(list1, list2) != nil {
		testCase.Errorf("Linked List Error: Linked list has no intersection but returned an intersection node")
	}

	intersect := NewLinkedListNode(4)
	intersect.InsertAtBack(5)
	intersect.InsertAtBack(6)

	list1.next.next.next = intersect
	list2.next.next.next = intersect

	node := FindIntersectionNode(list1, list2)
	if node == nil || node.data != 4 {
		testCase.Errorf("Linked List Error: Linked list has intersection but returned the wrong intersection node")
	}
}

// Test to detect a cycle in a linked list
func TestDetectCycleInLinkedList(testCase *testing.T) {
	testCase.Log("To test a cycle in a linked list")

	head := NewLinkedListNode(1)
	head.InsertAtBack(2)
	head.InsertAtBack(3)

	if DetectCycleInLinkedList(head) != nil {
		testCase.Errorf("Linked List Error: Linked list has no cycle but returned a cycle node")
	}

	intersect := NewLinkedListNode(4)
	intersect.InsertAtBack(5)
	intersect.InsertAtBack(6)

	head.next.next.next = intersect
	intersect.next.next = head.next.next

	node := DetectCycleInLinkedList(head)
	if node == nil || node.data != 3 {
		testCase.Errorf("Linked List Error: Linked list has a cycle but returned the wrong cycle node")
	}
}
