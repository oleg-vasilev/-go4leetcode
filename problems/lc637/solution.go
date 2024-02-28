package lc637

import (
	"container/list"
)

// 637. Average of Levels in Binary Tree

// Given the root of a binary tree, return the average value of the
// nodes on each level in the form of an array.
// Answers within 10-5 of the actual answer will be accepted.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var queue = list.New()
	queue.PushBack(root)

	var result = make([]float64, 0)

	var (
		count, sum int
	)
	for queue.Len() > 0 {
		count = queue.Len()
		sum = 0
		for i := 0; i < count; i++ {
			item := queue.Front()
			queue.Remove(item)
			node := item.Value.(*TreeNode)
			sum += node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, float64(sum)/float64(count))
	}
	return result
}
