package llists

import "sync"

// LinkedList is a single Node in the list
type LinkedList struct {
	lock sync.Mutex
	data int
	next *LinkedList
}

// DoublyLinkedList is a single Node in a doubly linked list
type DoublyLinkedList struct {
	lock sync.Mutex
	data int
	next *DoublyLinkedList
	prev *DoublyLinkedList
}