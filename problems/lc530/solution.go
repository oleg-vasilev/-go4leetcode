package lc530

import (
	"math"
)

// 530. Minimum Absolute Difference in BST

// Given the root of a Binary Search Tree (BST), return the minimum
// absolute difference between the values of any two different nodes in the tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var result = math.MaxInt
	var lastRootValue = math.MaxInt

	var inorderTraverse func(root *TreeNode)
	inorderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// left
		inorderTraverse(root.Left)
		// root
		result = min(result, abs(root.Val-lastRootValue))
		lastRootValue = root.Val
		// right
		inorderTraverse(root.Right)
	}
	inorderTraverse(root)
	return result
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getMinimumDifferenceBFS(root *TreeNode) int {
	var result = math.MaxInt

	var bfs func(root *TreeNode) (int, int)

	bfs = func(root *TreeNode) (int, int) {
		smallest, biggest := root.Val, root.Val

		if root.Left != nil {
			smallestL, biggestL := bfs(root.Left)

			result = min(result, root.Val-biggestL)
			smallest = smallestL
		}

		if root.Right != nil {
			smallestR, biggestR := bfs(root.Right)

			result = min(result, smallestR-root.Val)
			biggest = biggestR
		}

		return smallest, biggest
	}

	bfs(root)
	return result
}
