package lc114

// 114. Flatten Binary Tree to Linked List

// Given the root of a binary tree, flatten the tree into a "linked list":
//
// The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
// The "linked list" should be in the same order as a pre-order traversal of the binary tree.
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var stack = make([]*TreeNode, 0)

	stack = append(stack, root.Right)
	stack = append(stack, root.Left)

	var last = root
	last.Left = nil
	last.Right = nil

	var current *TreeNode
	for len(stack) > 0 {
		// pop top item
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current == nil { // skip nulls
			continue
		}
		// push childs to stack
		stack = append(stack, current.Right)
		stack = append(stack, current.Left)
		// link itself to the chain
		last.Right = current
		// clear links and move last to new item
		last = last.Right
		last.Left = nil
		last.Right = nil
	}
}
