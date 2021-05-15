package trees

import "testing"

// Test the in-order traversal of the binary tree
func TestInorderTraversalIterative(testCase *testing.T) {
	testCase.Log("To test the in-order traversal of the binary tree")

	var tree *Node = nil
	tree, _ = InsertBinarySearchTree(tree, 5)
	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	order := InorderTraversalIterative(tree)
	if len(order) != 5 {
		testCase.Errorf("Traversal Error: Order length does not match the number of nodes in the tree")
	}

	if order[0].data != 2 || order[1].data != 4 || order[2].data != 5 || order[3].data != 6 || order[4].data != 7 {
		testCase.Errorf("Traversal Error: Order in which values were returned not correct")
	}
}

// Test the pre-order traversal of the binary tree
func TestPreorderTraversalIterative(testCase *testing.T) {
	testCase.Log("To test the pre-order traversal of the binary tree")

	var tree *Node = nil
	tree, _ = InsertBinarySearchTree(tree, 5)
	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	order := PreorderTraversalIterative(tree)
	if len(order) != 5 {
		testCase.Errorf("Traversal Error: Order length does not match the number of nodes in the tree")
	}

	if order[0].data != 5 || order[1].data != 2 || order[2].data != 4 || order[3].data != 7 || order[4].data != 6 {
		testCase.Errorf("Traversal Error: Order in which values were returned not correct")
	}
}

// Test the post-order traversal of the binary tree
func TestPostorderTraversalIterative(testCase *testing.T) {
	testCase.Log("To test the post-order traversal of the binary tree")

	var tree *Node = nil
	tree, _ = InsertBinarySearchTree(tree, 5)
	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	order := PostorderTraversalIterative(tree)
	if len(order) != 5 {
		testCase.Errorf("Traversal Error: Order length does not match the number of nodes in the tree")
	}

	if order[0].data != 4 || order[1].data != 2 || order[2].data != 6 || order[3].data != 7 || order[4].data != 5 {
		testCase.Errorf("Traversal Error: Order in which values were returned not correct")
	}
}

// Test the level-order traversal of the binary tree
func TestLevelorderTraversalIterative(testCase *testing.T) {
	testCase.Log("To test the level-order traversal of the binary tree")

	var tree *Node = nil
	tree, _ = InsertBinarySearchTree(tree, 5)
	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	order := LevelorderTraversalIterative(tree)
	if len(order) != 5 {
		testCase.Errorf("Traversal Error: Order length does not match the number of nodes in the tree")
	}

	if order[0].data != 5 || order[1].data != 2 || order[2].data != 7 || order[3].data != 4 || order[4].data != 6 {
		testCase.Errorf("Traversal Error: Order in which values were returned not correct")
	}
}

// Test the level-order traversal of the binary tree
func TestLevelorderTraversalIterativeReturnsTwoDimensional(testCase *testing.T) {
	testCase.Log("To test the level-order traversal of the binary tree")

	var tree *Node = nil
	tree, _ = InsertBinarySearchTree(tree, 5)
	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	order := LevelorderTraversalIterativeReturnsTwoDimensional(tree)
	if len(order) != 3 || len(order[0]) != 1 || len(order[1]) != 2 || len(order[2]) != 2 {
		testCase.Errorf("Traversal Error: Order length does not match the number of nodes in the tree")
	}

	if order[0][0].data != 5 || order[1][0].data != 2 || order[1][1].data != 7 || order[2][0].data != 4 || order[2][1].data != 6 {
		testCase.Errorf("Traversal Error: Order in which values were returned not correct")
	}
}
