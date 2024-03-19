package lc82

// 82. Remove Duplicates from Sorted List II

// Given the head of a sorted linked list, delete all nodes that have
// duplicate numbers, leaving only distinct numbers from the original list.
// Return the linked list sorted as well.
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var fakeHead = new(ListNode)
	fakeHead.Next = head

	var prev = fakeHead
	var left = head
	var right = head.Next
	for right != nil {
		if left.Val != right.Val {
			// values are different - move all three pointers
			prev = left
			left = right
			right = right.Next
			continue
		}
		// left and right values the same and both should be skipped
		// including all possible duplicates - find next different value using right pointer
		for right.Next != nil && right.Val == right.Next.Val {
			right = right.Next
		}
		// left and right now indicate the boundaries of repeating elements that should be skipped
		// so new left should be next value after the current right pointer and prev should point to this new left
		left = right.Next
		prev.Next = left

		if left != nil {
			right = left.Next
		} else {
			break
		}
	}
	return fakeHead.Next
}
