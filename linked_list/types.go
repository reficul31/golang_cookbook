package llists

import "sync"

// LinkedList is a single Node in the list
type LinkedList struct {
	lock sync.Mutex
	data int
	next *LinkedList
}