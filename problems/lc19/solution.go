package lc19

/**
 * Remove Nth item from end of singly-linked list
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var preTarget *ListNode = head
	var end *ListNode = head
	var distance int
	for end.Next != nil {
		end = end.Next
		if distance == n {
			preTarget = preTarget.Next
		} else {
			distance++
		}
	}
	if distance < 1 {
		return nil
	}
	if distance < n {
		return head.Next
	}
	if preTarget.Next != nil {
		preTarget.Next = preTarget.Next.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
