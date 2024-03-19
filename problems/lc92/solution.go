package lc92

// 92. Reverse Linked List II

type ListNode struct {
	Val  int
	Next *ListNode
}

// One pass straightforward solution without any tricks
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if left == right {
		return head
	}
	// meaningful nodes we want to find
	var lastLeftUntouched *ListNode
	var firstLeftUntouched *ListNode
	var lastReversedNode *ListNode
	var firstReversedNode *ListNode

	// default variables for reversing linked list
	var curr, prev, next *ListNode
	curr = head
	for currIdx := 1; curr != nil; currIdx++ {
		if currIdx < left {
			if currIdx == left-1 {
				lastLeftUntouched = curr
			}
			curr = curr.Next
			continue
		}
		if currIdx == left {
			lastReversedNode = curr
		}
		if currIdx == right {
			firstReversedNode = curr
		}
		if currIdx == right+1 {
			firstLeftUntouched = curr
			break
		}
		// default reverse logic
		next, curr.Next = curr.Next, prev
		prev, curr = curr, next
	}
	// fix links between left, middle and right parts
	if lastLeftUntouched != nil {
		lastLeftUntouched.Next = firstReversedNode
	}
	if lastReversedNode != nil {
		lastReversedNode.Next = firstLeftUntouched
	}
	// check if head node was touched by reversing
	if left == 1 {
		return firstReversedNode
	}
	return head
}
