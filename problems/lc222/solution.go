package lc222

import (
	"math"
)

// 222. Count Complete Tree Nodes

// Given the root of a complete binary tree, return the number
// of the nodes in the tree.
//
// According to Wikipedia, every level, except possibly the last,
// is completely filled in a complete binary tree, and all nodes
// in the last level are as far left as possible.
// It can have between 1 and 2h nodes inclusive at the last level h.
//
// Design an algorithm that runs in less than O(n) time complexity.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// default O(n)
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
// 	return 1 + countNodes(root.Left) + countNodes(root.Right)
// }

// less than O(n), something ~O((logn)^2)
func countNodes(root *TreeNode) int {
	var left = leftHeight(root)
	var right = rightHeight(root)
	// perfect bintree
	if left == right {
		return int(math.Pow(2, float64(left)) - 1)
	}
	// not perfect - count each subtree individually basically until
	// we find two perfect trees with different heights
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func leftHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + leftHeight(root.Left)
}

func rightHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + rightHeight(root.Right)
}
