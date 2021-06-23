package llists

// Function to delete duplicates from a linked list using hash maps
// @arg head - Head of the linked list
// @returns - Modified linked list with the duplicates deleted
func DeleteDuplicateUsingHashMaps(head *LinkedList) *LinkedList {
	if head == nil {
		return nil
	}

	hashSet := make(map[int]bool)
	var previous *LinkedList = nil

	for current := head; current.next != nil; {
		if val, err := hashSet[current.next.data]; err && val {
			if previous != nil {
				previous.next = current.next
				current = current.next
			} else {
				head = head.next
				current = head
			}
		} else {
			hashSet[current.data] = true
			current = current.next
		}
	}
	return head
}

// Function to delete duplicates from a linkedlist using multiple iterations over the list
// @arg head - Head of the linked list
// @returns - Modified linked list with the duplicates deleted
func DeleteDuplicateUsingIteration(head *LinkedList) *LinkedList {
	for current := head; current != nil; {
		for runner := current; runner.next != nil; {
			if runner.next.data == current.data {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		current = current.next
	}
	return head
}

// Function to find the n nodes before the last node in a linked list
// @arg head - Head of the linked list
// @arg n - Number of nodes skipped before the end of the list
func nthToLastNode(head *LinkedList, n int) *LinkedList {
	if head == nil {
		return nil
	}

	current, runner := head, head
	for i := 0; i < n; i = i + 1 {
		if runner != nil {
			runner = runner.next
		}
	}

	for runner != nil {
		current = current.next
		runner = runner.next
	}

	return current
}

// Function to partition the linked list around the value n passed
// @arg head - Head of the linked list
// @arg n - Integer value around which the list needs to be partitioned
// @returns - The partitioned linked list around n
func PartitionAtN(head *LinkedList, n int) *LinkedList {
	if head == nil {
		return nil
	}

	var tail *LinkedList
	for tail = head; tail.next != nil; tail = tail.next {
	}

	node := head
	for node != nil {
		next := node.next
		if node.data > n {
			tail.next = node
			tail = node
		} else {
			node.next = head
			head = node
		}
		node = next
	}
	tail.next = nil
	return head
}

// Function to add two linked lists in position
// @arg head1 - Head of the linked list 1
// @arg head2 - Head of the linked list 2
// @returns - The final linked list with added values
func AddTwoNumbersUsingLinkedList(head1 *LinkedList, head2 *LinkedList) *LinkedList {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	carry := 0
	var result *LinkedList = nil
	trav1, trav2 := head1, head2
	trav3 := result

	for trav1 != nil || trav2 != nil || carry > 0 {
		sum1, sum2 := 0, 0
		if trav1 != nil {
			sum1 = trav1.data
			trav1 = trav1.next
		}

		if trav2 != nil {
			sum2 = trav2.data
			trav2 = trav2.next
		}

		sum := sum1 + sum2 + carry
		carry = sum / 10

		new := NewLinkedListNode(sum % 10)
		if result == nil {
			result = new
			trav3 = new
		} else {
			trav3.next = new
			trav3 = trav3.next
		}
	}

	return result
}
