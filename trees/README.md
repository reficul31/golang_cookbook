# Trees
Trees are hierarchical data structures. The topmost node of the tree is called the root. All the subsequent nodes attached to the node are called children.

## 1. Binary Tree
A tree having two children are called binary tree. 

### 1.1 Properties
The properties of the binary tree are as follow:  
1. Maximum number of nodes at level is pow(2, level).
2. Maximum number of nodes at a height is pow(2, height) - 1.
3. Minimum possible height of a binary tree having N nodes is Log2(N + 1).
4. A binary tree with L leaves has at least Log2(L) + 1 levels.

### 1.2 Types of Binary Trees
There are many types of binary trees and each of them have different uses. Following is the list of the most used binary trees.
1. **Binary Search Tree**: The binary search tree (BST) has the following properties:
	1. The left subtree of a node contains only nodes with keys lesser than the node’s key.
	2. The right subtree of a node contains only nodes with keys greater than the node’s key.
	3. The left and right subtree each must also be a binary search tree.
2. **Full Binary Tree**: A Binary Tree is a full binary tree if every node has 0 or 2 children. We can also say a full binary tree is a binary tree in which all nodes except leaf nodes have two children.
3. **Complete Binary Tree**:  A Binary Tree is a Complete Binary Tree if all the levels are completely filled except possibly the last level and the last level has all keys as left as possible.
4. **Perfect Binary Tree**: A Binary tree is a Perfect Binary Tree in which all the internal nodes have two children and all leaf nodes are at the same level.
5. **Balanced Binary Tree**: A binary tree is balanced if the height of the tree is O(Log n) where n is the number of nodes.
	1. **AVL Tree**: The tree maintains O(Log n) height by making sure that the difference between the heights of the left and right subtrees is at most 1.
	2. **Red Black Tree**: The trees maintain O(Log n) height by making sure that the number of Black nodes on every root to leaf paths is the same and there are no adjacent red nodes

### 1.3 Binary Tree Traversals
Traversal in trees mean the order in which the values are visited by the user. There are multiple ways to traverse a binary tree and each have their own different implementation.
#### 1.3.1 In-Order Traversal
In in-order traversal, we first visit the left node then the root node and then the right node.  

Algorithm InorderTraversal(tree):
1. Traverse the left subtree
2. Visit the root node
3. Traverse the right subtree

This algorithm can be implemented in multiple ways.
##### 1.3.1.1 In-Order Traversal Iterative
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
##### 1.3.1.2 In-Order Traversal Recursive
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L31)]  
```
Algorithm InorderTraversalRecursive(tree)
1. InorderTraversalRecursive(tree's left child)
2. Visit tree node
3. InorderTraversalRecursive(tree's right child)
```
#### 1.3.2 Pre-Order Traversal
In pre-order traversal, we first visit the root node then the left child and then the right child.
Algorithm PreorderTraversal(tree)
1. Visit the root node
2. Traverse the left subtree
3. Traverse the right subtree

This algorithm can also be implemented in multiple ways.
##### 1.3.2.1 Pre-Order Traversal Iterative
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
##### 1.3.2.2 Pre-Order Traversal Recursive
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L69)]  
```
Algorithm PreorderTraversalRecursive(tree):
1. Visit the tree node
2. PreorderTraversalRecursive(tree's left child)
3. PreorderTraversalRecursive(tree's right child)
```
#### 1.3.3 Post-Order Traversal
In post-order traversal, we first visit the left node of the tree then the right node of the tree and then finally the root node.
Algorithm PostorderTraversal(tree)
1. Traverse the left subtree
2. Traverse the right subtree
3. Visit the root node


This algorithm can also be implemented in multiple ways:
##### 1.3.3.1 Post-Order Traversal Iterative
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
##### 1.3.3.2 Post-Order Traversal Recursive
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/trees/traversal.go#L111)]
```
Algorithm PreorderTraversalRecursive(tree):
1. PreorderTraversalRecursive(tree's left child)
2. PreorderTraversalRecursive(tree's right child)
3. Visit the tree node
```
#### 1.3.4 Level-Order Traversal
In level-order traversal, we visit the nodes at the same level. The nodes that are on the same height are visited together and then we visit the nodes on the next level.  
Algorithm LevelOrderTraversal(tree)
1. Visit the node at the current level
2. Traverse to the next level of the tree


##### 1.3.4.1 Level-Order Traversal (Flat order returns)
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
##### 1.3.4.2 Level-Order Traversal (2 dimensional returns)
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