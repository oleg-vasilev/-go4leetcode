package lc103

import (
	"container/list"
	"slices"
)

// 103. Binary Tree Zigzag Level Order Traversal

// Given the root of a binary tree, return the zigzag level order
// traversal of its nodes' values. (i.e., from left to right, then
// right to left for the next level and alternate between).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Same default BFS but with additional tracking reverse order flag and reversing level result slice

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0)
	if root == nil {
		return res
	}

	var queue = list.New()
	var levelNodesCount int
	var reverseOrder bool
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelNodesCount = queue.Len()
		var levelNodes = make([]int, 0, levelNodesCount)
		for i := 0; i < levelNodesCount; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			levelNodes = append(levelNodes, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		if reverseOrder {
			slices.Reverse(levelNodes)
		}
		res = append(res, levelNodes)
		reverseOrder = !reverseOrder
	}
	return res
}
