package llists

import "sync"

// NewDoublyLinkedListNode is used to return a new node pointer of type DoublyLinkedList 
func NewDoublyLinkedListNode(element int) *DoublyLinkedList {
	return &DoublyLinkedList{
		sync.Mutex{},
		element,
		nil,
		nil,
	}
}

// InsertAtHeadDLL inserts a node at the head of the doubly linked list
func (head *DoublyLinkedList) InsertAtHeadDLL(element int) {
	head.lock.Lock()
	defer head.lock.Unlock()

	new := NewDoublyLinkedListNode(element)
	new.next = head.next
	new.prev = head
	if head.next != nil {
		head.next.prev = new
	}
	head.next = new
}

// InsertAtBackDLL inserts a node at the back of the doubly linked list
func (head *DoublyLinkedList) InsertAtBackDLL(element int) {
	head.lock.Lock()
	defer head.lock.Unlock()

	new := NewDoublyLinkedListNode(element)
	var temp *DoublyLinkedList
	for temp = head ; temp.next!=nil ; temp = temp.next {}
	temp.next = new
	new.prev = temp
}

// DeleteFromHeadDLL deletes a node from the head of a doubly linked list
func (head *DoublyLinkedList) DeleteFromHeadDLL() (int, error) {
	head.lock.Lock()
	defer head.lock.Unlock()

	if head.next == nil {
		return -1, ErrLinkedListEmpty
	}

	temp := head.next
	head.next = temp.next
	if temp.next != nil {
		temp.next.prev = head
	}
	return temp.data, nil
}

// DeleteFromBackDLL deletes a node from the back of a doubly linked list
func (head *DoublyLinkedList) DeleteFromBackDLL() (int, error) {
	head.lock.Lock()
	defer head.lock.Unlock()

	if head.next == nil {
		return -1, ErrLinkedListEmpty
	}

	var temp *DoublyLinkedList
	for temp=head ; temp.next!=nil ; temp=temp.next {}
	temp.prev.next = nil
	return temp.data, nil
}
