package lc876

// Mid of the list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var mid *ListNode = head
	var end *ListNode = head
	distance := 0 // distance between mid & end
	for end.Next != nil {
		end = end.Next
		distance++
		if distance == 2 || end.Next == nil {
			mid = mid.Next
			distance = 0
		}
	}
	return mid
}

func middleNodeV2(head *ListNode) *ListNode {
	var mid *ListNode = head
	var end *ListNode = head
	moveMid := false
	for end.Next != nil {
		end = end.Next
		moveMid = !moveMid
		if moveMid {
			mid = mid.Next
		}
	}
	return mid
}

type ListNode struct {
	Val  int
	Next *ListNode
}
