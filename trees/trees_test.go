package trees

import "testing"

var err error

// Test the new node memory allocation function of NewNode
func TestNewNode(testCase *testing.T) {
	testCase.Log("To test that the new tree node is allocated the correct value")

	node := NewNode(1)
	if node.data != 1 || node.left != nil || node.right != nil {
		testCase.Errorf("Tree Error: Node not allocated the correct value")
	}
}

// Test the tree construction using InsertBinarySearchTree
func TestInsertBinarySearchTree(testCase *testing.T) {
	testCase.Log("To test the construction of binary search tree using InsertBinarySearchTree")
	var tree *Node
	tree = nil

	if tree, err = InsertBinarySearchTree(tree, 5); err != nil {
		testCase.Errorf("Tree Error: %s", err)
	}
	if tree.data != 5 || tree.left != nil || tree.right != nil {
		testCase.Errorf("Tree Error: Node not allocated the correct value")
	}

	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	if tree.left.data != 2 || tree.right.data != 7 {
		testCase.Errorf("Tree Error: Binary search tree not created correctly")
	}
	if tree.left.right.data != 4 || tree.right.left.data != 6 {
		testCase.Errorf("Tree Error: Binary search tree not created correctly")
	}

	if tree, err = InsertBinarySearchTree(tree, 2); err == nil {
		testCase.Errorf("Tree Error: Duplicate Node error not reported")
	}
}
