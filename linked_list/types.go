package llists

import "sync"

type LinkedList struct {
	lock sync.Mutex
	data int
	next *LinkedList
}