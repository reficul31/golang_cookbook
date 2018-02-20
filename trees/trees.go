package trees

import "sync"

// NewNode is used to return a pointer of type Node
func NewNode(element int) *Node {
	return &Node{
		sync.Mutex{},
		element,
		nil,
		nil,
	}
}

// InsertBinarySearchTree is used to construct a binary search tree from an integer slice
func InsertBinarySearchTree(tree *Node, element int) (*Node, error) {
	if tree == nil {
		return NewNode(element), nil
	}

	if tree.data > element {
		if tree.left, err = InsertBinarySearchTree(tree.left, element) ; err != nil {
			return tree, err
		}
	} else if tree.data < element {
		if tree.right, err = InsertBinarySearchTree(tree.right, element) ; err != nil {
			return tree, err
		}
	} else {
		return tree, ErrDuplicateNode
	}
	return tree, nil
}	