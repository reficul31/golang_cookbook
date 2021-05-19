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