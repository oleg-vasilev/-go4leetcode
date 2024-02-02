package lc21

type ListNode struct {
	Val  int
	Next *ListNode
}

// Straightforward non-recursive solution
func mergeTwoListsNoRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	
	var resListFirst *ListNode
	var resNext *ListNode
	var resListLast, first, second *ListNode
	
	first = list1
	second = list2
	
	for first != nil || second != nil {
		if first == nil {
			resNext = second
			second = second.Next
		} else if second == nil {
			resNext = first
			first = first.Next
		} else {
			// both items represented
			if first.Val < second.Val {
				resNext = first
				first = first.Next
			} else {
				resNext = second
				second = second.Next
			}
		}
		// now resNext contains valid next item for result list
		if resListFirst == nil {
			resListFirst = resNext
			resListLast = resNext
		} else {
			resListLast.Next = resNext
			resListLast = resListLast.Next
		}
	}
	return resListFirst
}

// Default recursive solution (modifies input data)
func mergeTwoLists(first *ListNode, second *ListNode) *ListNode {
	var resNext *ListNode
	
	if first == nil && second == nil {
		return nil
	}
	
	if first == nil && second != nil {
		resNext = second
		second = second.Next
	} else if second == nil && first != nil {
		resNext = first
		first = first.Next
	} else {
		if first.Val < second.Val {
			resNext = first
			first = first.Next
		} else {
			resNext = second
			second = second.Next
		}
	}
	resNext.Next = mergeTwoLists(first, second)
	return resNext
}
