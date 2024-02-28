package lc2

// 2. Add Two Numbers

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultHead *ListNode
	var currentNode *ListNode

	var overflow int
	var currLeft = l1
	var currRight = l2

	for currLeft != nil || currRight != nil || overflow != 0 {
		// new node creation for result
		if currentNode == nil {
			currentNode = new(ListNode)
			resultHead = currentNode
		} else {
			currentNode.Next = new(ListNode)
			currentNode = currentNode.Next
		}
		// calculate sum
		currentNode.Val = overflow
		if currRight != nil {
			currentNode.Val += currRight.Val
			currRight = currRight.Next
		}
		if currLeft != nil {
			currentNode.Val += currLeft.Val
			currLeft = currLeft.Next
		}
		// fix overflow
		overflow = currentNode.Val / 10
		currentNode.Val = currentNode.Val % 10
	}
	return resultHead
}
