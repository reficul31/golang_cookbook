package llists

import (
	"sync"
	"errors"
)

// Function to allocate memory to a new node
func NewLinkedListNode(element int) *LinkedList {
	return &LinkedList{
		sync.Mutex{},
		element,
		nil,
	}
}

// Function to insert a node at the head of the list
func (head *LinkedList) InsertAtHead(element int) {
	head.lock.Lock()
	defer head.lock.Unlock()

	new := NewLinkedListNode(element)
	new.next = head.next
	head.next = new
}

// Function to insert a node at the back of the list
func (head *LinkedList) InsertAtBack(element int) {
	head.lock.Lock()
	defer head.lock.Unlock()

	var temp *LinkedList
	for temp = head ; temp.next!=nil ; temp = temp.next {}
	temp.next = NewLinkedListNode(element)
}

// Function to delete a node from the head of the list
func (head *LinkedList) DeleteFromHead() (int, error) {
	head.lock.Lock()
	defer head.lock.Unlock()

	if head.next == nil {
		return -1, errors.New("Linked List Empty")
	} else {
		temp := head.next
		head.next = temp.next
		return temp.data, nil
	}
}

// Function to delete a node from the back of the list
func (head *LinkedList) DeleteFromBack() (int, error) {
	head.lock.Lock()
	defer head.lock.Unlock()

	if head.next == nil {
		return -1, errors.New("Linked List Empty")
	} else {
		var prev, temp *LinkedList
		prev = head
		temp = head.next
		for temp.next !=nil {
			prev = temp
			temp = temp.next
		}
		prev.next = nil
		return temp.data, nil
	}
}