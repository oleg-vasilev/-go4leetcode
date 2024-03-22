package lc61

// 61. Rotate List

// Given the head of a linked list, rotate the list to the right by k places.

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	// find out list length, and link last item to first
	var curr *ListNode
	var length int
	curr = head
	length++
	for curr.Next != nil {
		curr = curr.Next
		length++
	}
	curr.Next = head
	// now we know length and list is circle
	var counterclockwiseSteps = k % length
	var clockwiseSteps = length - counterclockwiseSteps

	curr = head
	// move head clockwise
	for clockwiseSteps > 0 {
		if clockwiseSteps == 1 {
			curr.Next, curr = nil, curr.Next
		} else {
			curr = curr.Next
		}
		clockwiseSteps--
	}
	return curr
}
