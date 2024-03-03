package lc102

import (
	"container/list"
)

// 102. Binary Tree Level Order Traversal

// Given the root of a binary tree, return the level order traversal
// of its nodes' values. (i.e., from left to right, level by level).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// default BFS

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0)
	if root == nil {
		return res
	}

	var queue = list.New()
	var levelNodesCount int
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
		res = append(res, levelNodes)
	}
	return res
}
