package trees

import "fmt"

// Iterative function for In-Order Traversal of a Binary Tree
// @arg root - Root of the tree node
// @returns - An array with the elements in the order of the traversal
func InorderTraversalIterative(root *Node) []*Node {
	if root == nil {
		return []*Node{}
	}

	stack, order := []*Node{}, []*Node{}
	current := root

	for current != nil || len(stack) != 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.left
		} else {
			current, stack = stack[len(stack)-1], stack[:len(stack)-1]
			order = append(order, current)
			current = current.right
		}
	}
	return order
}

// Recursive function for In-Order Traversal of a Binary Tree
// @arg root - Root of the tree node
func InorderTraversalRecursive(root *Node) {
	if root == nil {
		return
	}

	InorderTraversalRecursive(root.left)
	fmt.Println(root.data)
	InorderTraversalRecursive(root.right)
}

// Iterative function for Pre-Order Traversal of a Binary Tree
// @arg root - Root of the tree node
// @returns - An array with the elements in the order of the traversal
func PreorderTraversalIterative(root *Node) []*Node {
	if root == nil {
		return []*Node{}
	}

	stack, order := []*Node{}, []*Node{}
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, node)

		if node.right != nil {
			stack = append(stack, node.right)
		}
		if node.left != nil {
			stack = append(stack, node.left)
		}
	}
	return order
}

// Recursive function for Pre-Order Traversal of a Binary Tree
// @arg root - Root of the tree node
func PreorderTraversalRecursive(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	PreorderTraversalRecursive(root.left)
	PreorderTraversalRecursive(root.right)
}

// Iterative function for Post-Order Traversal of a Binary Tree
// @arg root - Root of the tree node
// @returns - An array with the elements in the order of the traversal
func PostorderTraversalIterative(root *Node) []*Node {
	if root == nil {
		return []*Node{}
	}

	stack, order := []*Node{}, []*Node{}
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, node)

		if node.left != nil {
			stack = append(stack, node.left)
		}
		if node.right != nil {
			stack = append(stack, node.right)
		}
	}

	for i, j := 0, len(order)-1; i <= j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order
}

// Recursive function for Post-Order Traversal of a Binary Tree
// @arg root - Root of the tree node
func PostorderTraversalRecursive(root *Node) {
	if root == nil {
		return
	}

	PostorderTraversalRecursive(root.left)
	PostorderTraversalRecursive(root.right)
	fmt.Println(root.data)
}

// Iterative function for Level Order traversal of a Binary Tree
// @arg root - Root of the tree node
// @returns - A one dimensional array with the elements in the order of the traversal
func LevelorderTraversalIterative(root *Node) []*Node {
	if root == nil {
		return []*Node{}
	}

	queue, order := []*Node{}, []*Node{}
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return order
}

// Iterative function for Level Order traversal of a Binary Tree
// @arg root - Root of the tree node
// @returns - A two dimensional array with the elements in the order of the traversal
func LevelorderTraversalIterativeReturnsTwoDimensional(root *Node) [][]*Node {
	if root == nil {
		return [][]*Node{}
	}

	queue, n_queue := []*Node{}, []*Node{}
	order, level := [][]*Node{}, []*Node{}

	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		level = append(level, node)

		if node.left != nil {
			n_queue = append(n_queue, node.left)
		}

		if node.right != nil {
			n_queue = append(n_queue, node.right)
		}

		if len(queue) == 1 {
			queue, n_queue = n_queue, []*Node{}
			order, level = append(order, level), []*Node{}
		} else {
			queue = queue[1:]
		}
	}
	return order
}
