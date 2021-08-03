package trees

import (
	"math"
)

// Function to get the maximum depth of the binary tree
// @arg root - Root node of the tree
// @returns - Depth of the tree
func DepthOfTree(root *Node) int {
	if root == nil {
		return 0
	}

	leftDepth := DepthOfTree(root.left)
	rightDept := DepthOfTree(root.right)

	return int(math.Max(float64(leftDepth), float64(rightDept))) + 1
}

// Function to check whether the binary tree is symmetric around its center
// @arg left - Left side of the tree
// @arg right - Right side of the tree
// @returns - Boolean value of whether the symmetric or not
func IsSymmetric(left, right *Node) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil && left.data == right.data {
		return IsSymmetric(left.left, right.right) && IsSymmetric(left.right, right.left)
	}

	return false
}

// Function to check whether a tree has a root-to-leaf path such that adding up all the values along the path equals targetSum
// @arg root - Root of the tree
// @arg sum - Target sum for the root to leaf path
// @returns - Boolean value of whether the root-to-leaf path contains the target sum
func HasTargetSum(root *Node, sum int) bool {
	if root == nil {
		return sum == 0
	}

	// Check right here if the node is a leaf node and the sum is 0
	target := sum - root.data
	if target == 0 && root.left == nil && root.right == nil {
		return true
	}

	ans := false
	if root.left != nil {
		ans = ans || HasTargetSum(root.left, target)
	}

	if root.right != nil {
		ans = ans || HasTargetSum(root.right, target)
	}

	return ans
}
