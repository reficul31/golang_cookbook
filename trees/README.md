# Tree Algorithms
## 1. Binary Tree Traversals
Traversal in trees mean the order in which the values are visited by the user. There are multiple ways to traverse a binary tree and each have their own different implementation.
### 1.1 In-Order Traversal
In in-order traversal, we first visit the left node then the root node and then the right node.  

Algorithm InorderTraversal(tree):
1. Traverse the left subtree
2. Visit the root node
3. Traverse the right subtree

This algorithm can be implemented in multiple ways.
#### 1.1.1 In-Order Traversal Iterative
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L8)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal_test.go#L6)]  
```
Algorithm InorderTraversalIterative(tree):
1. Initialize an empty stack
2. Set current to the root node of the tree
3. Do the following till current != null and stack is not empty
	1. If current is not null
		1. Push current on the stack
		2. Set current to current's left child
    2. If current is null
    	1. Set current to the top of stack
    	2. Pop from the stack
    	3. Visit current node
    	4. Set current to current's right child
```
#### 1.1.2 In-Order Traversal Recursive
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L31)]  
```
Algorithm InorderTraversalRecursive(tree)
1. InorderTraversalRecursive(tree's left child)
2. Visit tree node
3. InorderTraversalRecursive(tree's right child)
```
### 1.2 Pre-Order Traversal
In pre-order traversal, we first visit the root node then the left child and then the right child.
Algorithm PreorderTraversal(tree)
1. Visit the root node
2. Traverse the left subtree
3. Traverse the right subtree

This algorithm can also be implemented in multiple ways.
#### 1.2.1 Pre-Order Traversal Iterative
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L44)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal_test.go#L27)]  
```
Algorithm PreorderTraversalIterative(tree):
1. Initialize an empty stack
2. Push the root node on top of stack
3. Do the following till stack is not empty
	1. Pop element from stack
	2. Visit the element
	3. Push right child of element on stack
	4. Push left child of element on stack
```
#### 1.2.2 Pre-Order Traversal Recursive
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L69)]  
```
Algorithm PreorderTraversalRecursive(tree):
1. Visit the tree node
2. PreorderTraversalRecursive(tree's left child)
3. PreorderTraversalRecursive(tree's right child)
```
### 1.3 Post-Order Traversal
In post-order traversal, we first visit the left node of the tree then the right node of the tree and then finally the root node.
Algorithm PostorderTraversal(tree)
1. Traverse the left subtree
2. Traverse the right subtree
3. Visit the root node


This algorithm can also be implemented in multiple ways:
#### 1.3.1 Post-Order Traversal Iterative
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L82)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal_test.go#L48)]  
```
Algorithm PostorderTraversalIterative(tree):
1. Initialize two empty stack
2. Push the root node to the top of stack 1
3. Do the following till stack 1 is not empty:
	1. Pop an element from stack 1
	2. Push the popped element into stack 2
	3. Push element's right child into stack 1
	4. Push element's left child into stack 1
4. Reverse stack 2 and visit nodes in order
```
#### 1.3.2 Post-Order Traversal Recursive
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L111)]
```
Algorithm PreorderTraversalRecursive(tree):
1. PreorderTraversalRecursive(tree's left child)
2. PreorderTraversalRecursive(tree's right child)
3. Visit the tree node
```
### 1.4 Level-Order Traversal
In level-order traversal, we visit the nodes at the same level. The nodes that are on the same height are visited together and then we visit the nodes on the next level.  
Algorithm LevelOrderTraversal(tree)
1. Visit the node at the current level
2. Traverse to the next level of the tree


#### 1.4.1 Level-Order Traversal (Flat order returns)
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L124)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal_test.go#L69)]  
```
Algorithm LevelOrderTraversalIterative(tree):
1. Initialize an empty queue
2. Push the root of the tree in the queue
3. Do the following till queue is not empty:
	1. Pop and element from front of queue
	2. Visit the popped element's data
	3. Push the element's left child into the queue
	4. Push the element's right child into the queue
```
#### 1.4.2 Level-Order Traversal (2 dimensional returns)
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L150)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal_test.go#L90)]  
```
Algorithm LevelOrderTraversalIterative(tree):
1. Initialize 2 empty queue
2. Push the root of the tree in the first queue
3. Do the following till the queue is not empty
	1. Pop the element from the front of queue 1
	2. Visit the popped element's data
	3. Push the element's left child in queue 2
	4. Push the element's right child in queue 2
	5. If queue 1 is empty
		1. Set queue 1 to queue 2
		2. Reinitialize queue 2 to an empty queue
		3. Do the operation after the level has been traversed
```
## 2. Recursive Problems
Recursion is one of the most commonly used techniques to tackle binary tree problems. The key to thinking in recursion is to think of the base case first and then taking the small tree and solving the problem for it. Then think how the solution for the small tree fits into the entire solution at hand. Following are some examples to illustrate the same.
### 2.1 Depth of Binary Tree
The depth of the binary tree is defined as the maximum number of nodes that exist between the root and the leaf node in a binary tree.
#### 2.1.1 Depth of Binary Tree Recursive
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems.go#L10)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems_test.go#L6)]  
```
Algorithm DepthOfTree(tree):
1. If tree == null, return 0
2. Set leftHeight as DepthOfTree(tree.Left)
3. Set rightHeight as DepthOfTree(tree.Right)
4. Return the maximum of leftHeight and rightHeight
```
#### 2.1.2 Depth of Binary Tree Iterative
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems.go#L24)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems_test.go#L23)]  
```
Algorithm DepthOfTreeIterative(tree):
1. Initialize 2 queue named queue and n_queue
2. Append the root node to the queue
3. While queue is not empty, do the following:
	1. Remove from queue and set to current
	2. If current's left is not null, insert current's left into n_queue
	3. If current's right is not null, insert current's right into n_queue
	4. If queue is empty, set queue to n_queue and initialize n_queue as a new queue and increment the depth
4. Return the depth
```
### 2.2 Symmetric Tree
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems.go#L57)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems_test.go#L40)]  
A symmetric tree is the one which has matching nodes on the left and right side of the root. Imagine putting a mirror at the node, the left subtree should resemble exactly the right subtree.
```
Algorithm IsSymmetric(left, right):
1. If left and right both are null return true
2. If either right or left is not null but the other is, return false
3. If right's data is not equal to left's data return false
4. Check if left's left and right's right are symmetric by calling the IsSymmetric function
5. Check if left's right and right's left are symmetric by calling the IsSymmetric function
6. Return the and of the results from step 4 and 5.
```
### 2.3 Target Sum
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems.go#L73)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/trees/problems_test.go#L64)]  
Given a target sum, we need to find a root to leaf node path in the tree such that the sum of the values in the path is equal to the target sum.
 
```
Algorithm HasTargetSum(root, targetSum):
1. If root is nil and sum is 0 return true, else return false
2. Set new sum as targetSum - root's data
3. If root's left is not null call HasTargetSum with root's left and new sum
4. If root's right is not null call HasTargetSum with root's right and new sum
5. Return the or of the results from step 4 an 5.
```