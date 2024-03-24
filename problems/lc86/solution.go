package lc86

// 86. Partition List

// Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
//
// You should preserve the original relative order of the nodes in each of the two partitions.

type ListNode struct {
	Val  int
	Next *ListNode
}

// straightforward with allocations of new list and without dummy heads trick
func partition(head *ListNode, x int) *ListNode {
	var lesserHead, lesserPtr, greaterPtr, greaterHead *ListNode

	var cur *ListNode
	cur = head
	for cur != nil {
		if cur.Val < x {
			if lesserPtr == nil {
				lesserPtr = new(ListNode)
				lesserPtr.Val = cur.Val
				lesserHead = lesserPtr
			} else {
				lesserPtr.Next = new(ListNode)
				lesserPtr.Next.Val = cur.Val
				lesserPtr = lesserPtr.Next
			}
		} else {
			if greaterPtr == nil {
				greaterPtr = new(ListNode)
				greaterPtr.Val = cur.Val
				greaterHead = greaterPtr
			} else {
				greaterPtr.Next = new(ListNode)
				greaterPtr.Next.Val = cur.Val
				greaterPtr = greaterPtr.Next
			}
		}
		cur = cur.Next
	}

	if lesserPtr != nil {
		lesserPtr.Next = greaterHead
		return lesserHead
	} else {
		return greaterHead
	}
}

func partition2(head *ListNode, x int) *ListNode {
	// dummy heads for two parts of the list
	var leftHead, rightHead = new(ListNode), new(ListNode)

	var leftPtr = leftHead
	var rightPtr = rightHead

	for head != nil {
		if head.Val < x {
			leftPtr.Next = head
			leftPtr = leftPtr.Next
		} else {
			rightPtr.Next = head
			rightPtr = rightPtr.Next
		}
		head = head.Next
	}
	rightPtr.Next = nil
	leftPtr.Next = rightHead.Next
	return leftHead.Next
}
