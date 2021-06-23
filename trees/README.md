# Tree Algorithms
## 1 Binary Tree Traversals
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
3. Do the following till current == null and stack is not empty
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
2. Push the rppt node to the top of stack 1
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