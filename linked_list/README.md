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

## 3. Linked List Algorithms
There are many problems that can be solved using linked lists. Here is a small subset of the problems that I have tackled in this project.

### 3.1 Delete Duplicates from Linked List
Given a linked list we need to delete all the duplicate nodes from the list. There are multiple ways to tackle this problem. Everything depends on the type of solution your are looking for.
#### 3.1.1 Using Hash Maps
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L6)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L8)]  
```
Algorithm DeleteDuplicates(head):
1. Initialize a hash map
2. Set current and previous to head which is the first element in the Linked List
3. While current != null, do the following
    1. If hashmap has a boolean value and it is true for the current data node
        1. Assign current to temp
        2. Set previous next and current to to current's next
        3. Free the memory space of temp
    2. If hashmap has a boolean value and it is false for the current data node
        1. Set hashmap is true for the current data node
        2. Set current to current's next
4. Return head which is the modified linked list
```

#### 3.1.2 Using Iteration
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L35)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L25)]  
```
Algorithm DeleteDuplicates(head):
1. Set current to head which is the first element in the linked list
2. While current != null, do the following
    1. Set runner to current
    2. While runner != null and runner's next != null, do the following
        1. If runner's next data is equal to current data
            1. Set runner's next to runner's next's next
            2. Free memory space for runner's next
        2. Set runner to runner's next
    3. Set current to current's next
3. Return head which is the modified linked list
```
### 3.2 Nth to the Last Node
Given a linked list we need to find the nth to the last node in the list.
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L52)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L42)]  
```
Algorithm NthToTheLastNode(head, n):
1. Set fast to head, which is the first element in the linked list
2. For i is equal to 0 till n, do the following:
    1. Set fast to fast's next
3. Set slow to head, which is the first element in the linked list
4. While fast != null, do the following
    1. Set fast to fast's next
    2. Set slow to slow's next
5. Return slow, which is the nth to the last node
```
### 3.3 Partition at N
Given a linked list and an integer value N. We need to return a modified linked list which is partitioned along N where elements on the left are lesser than N and element to the right are greater than N.
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L76)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L59)]  
```
Algorithm PartitionAtN(head, N):
1. Set new head and new tail to head, which is the first element in the linked list
2. Set traversal point to head
3. Do the following while traversal != null:
    1. Set next equal to traversal's next
    2. If traversal's data is less than N
        1. Set traversal's next to new head
        2. Set new head to traversal
    3. If traversal's data is greater than or equal to N:
        1. Set tail's next to traversal
        2. Set tail to traversal
        3. Set tail's next to null
    4. Set traversal to next
4. Return new head as the modified linked list
```
### 3.4 Add Numbers using Linked List
Given two numbers in reverse order stored as linked list. Return the addition of the two numbers as a new linked list.
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L103)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L76)]  
```
Algorithm AddNumbersUsingLinkedList(number1, number2):
1. Set traversal 1 to number 1 and traversal 2 to number 2
2. Create a new linked list named result
3. Set carry integer to 0
4. While traversal 1 or traversal 2's next != null or carry > 0:
    1. Set additive1 to traversal 1's data if traversal 1 not null
    2. Set additive2 to traversal 2's data if traversal 2 not null
    3. Set sum equal to additive 1 + additive 2 + carry
    4. Create a new node in the result list with data as sum%10
    5. Set carry as sum/10
5. Return result as the new list
```
### 3.5 Palindrome
Given a linked list we need to return whether the list is a palindrome or not.
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L148)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L96)]  
```
Algorithm IsPalindrome(head):
1. Set slow and fast to head which is the first element in the list
2. Initialize an empty stack
3. While fast != null and fast's next != null, do the following:
    1. Push the slow into the stack
    2. Set slow to slow's next
    3. Set fast to fast's next's next
4. If fast is equal to nil:
    1. Set slow to slow's next
5. While slow != null, do the following:
    1. Pop from the stack into node
    2. If node's data and slow's data dont match
        1. Return false as list is not a palindrome
    3. Set slow to slow's next
6. Return true as the list is a palindrome
```
### 3.6 Find Intersection Node
Given a two linked lists. We need to return the node at which the two lists intersect.
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L182)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L121)]  
```
Algorithm FindIntersectionNode(head1, head2):
1. Calculate the length of both linked lists
2. Set traversal 1 to head 1 and traversal 2 to head 2
2. If length 1 > length2:
    1. While i = 0 and i < length 1 - length 2 and i++, do:
        1. Set traversal 1 to traversal 1's next
3. If length 2 > length1:
    1. While i = 0 and i < length 2 - length 1 and i++, do:
        1. Set traversal 2 to traversal 2's next
4. While traversal 1 != null and traversal 2 != null, do the following:
    1. If traversal 1 == traversal 2:
        1. Return traversal 1 as the intersection node
    2. Set traversal 1 as traversal 1's next
    3. Set traversal 2 as traversal 2's next
5. Return null as no intersection was found
```
### 3.7 Find Cycle Node in Linked List
Given a linked list node, we need to find whether it contains a cycle and if it does we need to return the start of the cycle node.
[[Code](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L227)] | [[Test](https://github.com/reficul31/golang_cookbook/blob/master/linked_list/problems.go#L150)]  
```
Algorithm FindCycleNode(head):
1. Set slow and fast to head which is the first element in the linked list
2. While fast != null and fast's next != null, do the following:
    1. Set slow to slow's next
    2. Set fast to fast's next
    3. If slow == fast, break
3. If fast == null, return null since the list doesn't contain a cycle
4. Set slow to head
5. While slow != fast, do the following:
    1. Set slow to slow's next
    2. Set fast to fast's next
6. Return slow/fast as the cyclic node
```
