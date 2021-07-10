# Linked List Algorithms
## 1. Singly Linked Lists
The singly linked list is a linked list which only references the node ahead of it. There are 4 different basic operations in a linked list.
### 1.1 Insert at head
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists.go#L18)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists_test.go#L16)]  
```
Algorithm InsertAtHead(element):
1. Create a new node on element
2. Assign the new node's next to head
3. Set the head as the new node
```
### 1.2 Insert at back
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists.go#L29)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists_test.go#L32)]  
```
Algorithm InsertAtBack(element):
1. Traverse to the tail of the linked list
2. Create a new node on element
3. Set the tail's next to the new node
```
### 1.3 Delete from head
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists.go#L41)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists_test.go#L48)]  
```
Algorithm DeleteFromHead():
1. Deleted element is head
2. Set head equal to head's next
3. Free memory space of the deleted element
```
### 1.4 Delete from back
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists.go#L56)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/llists_test.go#L70)]  
```
Algorithm DeleteFromBack():
1. Traverse to the tail of the linked list while maintaining a previous node pointer
2. Set the previous node's next to nil
3. Free memory space of the tail element
```
# 2. Doubly Linked List
The singly linked list is a linked list which references not only the node ahead of it but the node before it as well. There are 4 different basic operations in a doubly linked list.
### 2.1 Insert at head
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists.go#L16)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists_test.go#L6)]  
```
Algorithm InsertAtHead(element):
1. Create a new node on element
2. If head is not null:
    1. Set new node's next and previous as null
    2. Set it as head and return
3. Else if head is null:
    1. Set new node's next as head and previous as null
    2. Set head's previous as new node
    3. Set the new node as head and return
```
### 2.2 Insert at back
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists.go#L30)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists_test.go#L37)]  
```
Algorithm InsertAtBack(element):
1. Traverse to the tail of the linked list
2. Create a new node on element
3. Set the tail's next to the new node
4. Set the new node's previous as tail
```
### 2.3 Delete from head
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists.go#L42)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists_test.go#54)]  
```
Algorithm DeleteFromHead():
1. Deleted element is head
2. Set head equal to head's next
3. Set head's previous as null
3. Free memory space of the deleted element
```
### 2.4 Delete from back
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists.go#L59)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/doubly_llists_test.go#L79)]  
```
Algorithm DeleteFromBack():
1. Traverse to the tail of the linked list while maintaining a previous node pointer
2. Set the previous node's next to nil
3. Free memory space of the tail element
```