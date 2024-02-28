package lc104

import (
	"container/list"
)

// 104. Maximum Depth of Binary Tree

// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest
// path from the root node down to the farthest leaf node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth = 0

	// queue with single element - root node (depth - 0)
	queue := list.New()
	queue.PushBack(root)

	// process three depths one by one
	for queue.Len() > 0 {
		// at this point queue contains only nodes from current depth level
		// remember nodes count to correctly iterate over them because in process
		// we will add more nodes to queue that wi do not want to touch in current iteration
		currentDepthNodesCount := queue.Len()
		// process each node at current "depth level" of tree
		// to add to queue all their childrens
		for i := 0; i < currentDepthNodesCount; i++ {
			// dequeue node
			nextElem := queue.Front()
			queue.Remove(nextElem)
			node := nextElem.Value.(*TreeNode)
			// add node childs to queue from the back
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		// at this point all nodes at current depth processed
		// and queue contains all childs - nodes at next depth
		// increase depth counter
		depth++
	}

	return depth
}
