package trees

import "sync"

// Node is the structure defining a single node of the tree
type Node struct {
	sync.Mutex
	data int
	left *Node
	right *Node
}