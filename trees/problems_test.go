package trees

import "testing"

// Test the depth of the binary tree function
func TestDepthOfTree(testCase *testing.T) {
	testCase.Log("To test the depth of the binary tree")

	var tree *Node = nil
	tree, _ = InsertBinarySearchTree(tree, 5)
	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	depth := DepthOfTree(tree)
	if depth != 3 {
		testCase.Errorf("Problems Error: Expected depth of tree to be %d. Instead received %d", 2, depth)
	}
}

// Test the symmetric function of a binary tree
func TestIsSymmetric(testCase *testing.T) {
	testCase.Log("To test the symmetric nature of the binary tree")

	var tree *Node = NewNode(5)
	tree.left = NewNode(3)
	tree.right = NewNode(3)
	tree.left.left = NewNode(4)
	tree.right.right = NewNode(4)
	tree.left.right = NewNode(5)
	tree.right.left = NewNode(6)

	symmetric := IsSymmetric(tree.left, tree.right)
	if symmetric {
		testCase.Errorf("Problems Error: Expected the binary tree not to be symmetric but was returned symmetric")
	}

	tree.right.left = NewNode(5)
	symmetric = IsSymmetric(tree.left, tree.right)
	if !symmetric {
		testCase.Errorf("Problems Error: Expected the binary tree to be symmetric but was returned not symmetric")
	}
}

// Test the function to check whether a root-to-leaf path contains the target sum
func TestHasTargetSum(testCase *testing.T) {
	testCase.Log("To test whether a root-to-leaf path contains the target sum")

	var tree *Node = nil
	tree, _ = InsertBinarySearchTree(tree, 5)
	tree, _ = InsertBinarySearchTree(tree, 2)
	tree, _ = InsertBinarySearchTree(tree, 4)
	tree, _ = InsertBinarySearchTree(tree, 7)
	tree, _ = InsertBinarySearchTree(tree, 6)

	hasSum := HasTargetSum(tree, 10)
	if hasSum {
		testCase.Errorf("Problems Error: Expected the target sum to not be present but was returned to be present")
	}

	hasSum = HasTargetSum(tree, 11)
	if !hasSum {
		testCase.Errorf("Problems Error: Expected the target sum to be present but was returned not to be present")
	}
}
