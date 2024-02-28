package lc141

// 141. Linked List Cycle

// Given head, the head of a linked list,
// determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the
// list that can be reached again by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next
// pointer is connected to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.

type ListNode struct {
	Val  int
	Next *ListNode
}

// Default hashmap solution
func hasCycle(head *ListNode) bool {
	var visited = make(map[*ListNode]bool)
	var current = head
	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}
	return false
}

// Beautiful memory-effective it cost of time
func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// fast moves twice faster
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			// we have a loop in they met
			return true
		}
	}
	// fast reached the end
	return false
}
