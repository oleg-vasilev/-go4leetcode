package lc206

// 206 - Reverse linked list

// Given the head of a singly linked list, reverse the list, and return the reversed list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var curr, prev, next *ListNode
	curr = head
	for curr != nil {
		// next = curr.Next
		// curr.Next = prev
		next, curr.Next = curr.Next, prev
		prev, curr = curr, next
	}
	return prev
}

// Recursive
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	tail := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return tail
// }
