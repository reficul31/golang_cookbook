package llists

import "sync"

// NewLinkedListNode is used to return a new node pointer of type LinkedList
func NewLinkedListNode(element int) *LinkedList {
	return &LinkedList{
		sync.Mutex{},
		element,
		nil,
	}
}

// InsertAtHead inserts a node at the head of the list
func (head *LinkedList) InsertAtHead(element int) {
	head.lock.Lock()
	defer head.lock.Unlock()

	new := NewLinkedListNode(element)
	new.next = head.next
	head.next = new
}

// InsertAtBack inserts a node at the back of the list
func (head *LinkedList) InsertAtBack(element int) {
	head.lock.Lock()
	defer head.lock.Unlock()

	var temp *LinkedList
	for temp = head ; temp.next!=nil ; temp = temp.next {}
	temp.next = NewLinkedListNode(element)
}

// DeleteFromHead deletes a node from the head of the list
func (head *LinkedList) DeleteFromHead() (int, error) {
	head.lock.Lock()
	defer head.lock.Unlock()

	if head.next == nil {
		return -1, ErrLinkedListEmpty
	}
	
	temp := head.next
	head.next = temp.next
	return temp.data, nil
}

// DeleteFromBack deletes a node from the back of the list
func (head *LinkedList) DeleteFromBack() (int, error) {
	head.lock.Lock()
	defer head.lock.Unlock()

	if head.next == nil {
		return -1, ErrLinkedListEmpty
	}
	
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