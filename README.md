# GoLang Cookbook
This repository is a collection of my implementations of common data structures and algorithms in GoLang. The main aim of this repository is to learn about GoLang in greater depth and in the process brush up on my Core Computer Science concepts.

# Linked Lists
A linked list is a linear data structure consisting of nodes. In a linked list node, each node is comprised of the data that it holds and an access to the next node in the data structure. Like arrays, Linked List is a linear data structure. Unlike arrays, linked list elements are not stored at a contiguous location; the elements are linked using pointers.

### Types of Linked Lists
1. Singly Linked List: In this type of linked list, the node contains a data member and an accessor to the next node in the list.
2. Doubly Linked List: In this type of linked list, the node contains a data member and an accessor to the next and the previous node in the list.
3. Circular Linked List: This type of linked list is unique since it can be made using DLLs and SLLs. Usually the last node in the linked list contains a null pointer reference denoting the end of the linked list. But in the circular linked list, there is no "end". The circular linked list forms a cycle where the reference of the last node in the list connects back to the first node in the list.

### Linked Lists v/s Arrays
#### Advantages
1. The size of arrays are fixed and therefore we must know the size of the data we are storing in advance.
2. Insertion and deletion can be expensive in an array since we need to create room in an array and then copy the elements to the new array.
#### Disadvantages
1. Random access of elements is not allowed. We have to access each element in the list sequentially.
2. Extra memory space for the pointer is required in addition to the data member.
3. Linked lists are not cache friendly. Since arrays are stored in contiguous memory location it makes it easy for the cache system.

### Representation
A linked list is represented by a pointer to the first node of the linked list. The first node is called the head. If the linked list is empty, then the value of the head is NULL. 
Each node in a list consists of at least two parts: 
1) Data
2) Pointer (Or Reference) to the next node 

# Trees
Trees are hierarchical data structures. The topmost node of the tree is called the root. All the subsequent nodes attached to the node are called children.

### 1. Binary Tree
A tree having two children are called binary tree. 

#### 1.1 Properties
The properties of the binary tree are as follow:  
1. Maximum number of nodes at level is pow(2, level).
2. Maximum number of nodes at a height is pow(2, height) - 1.
3. Minimum possible height of a binary tree having N nodes is Log2(N + 1).
4. A binary tree with L leaves has at least Log2(L) + 1 levels.

#### 1.2 Types of Binary Trees
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
