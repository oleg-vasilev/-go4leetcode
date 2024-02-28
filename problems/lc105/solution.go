package lc105

import (
	"slices"
)

// 105. Construct Binary Tree from Preorder and Inorder Traversal

// Given two integer arrays preorder and inorder where preorder is the
// preorder traversal of a binary tree and inorder is the inorder traversal
// of the same tree, construct and return the binary tree.
//

// Thoughts:
// - preorder has the node first. But you don't know the size of either branch.
// - inorder ALWAYS has the left branch to the left of the node, and right branch right of the node. So now you know the size of each branch.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// base case #1
	if len(preorder) == 0 {
		return nil
	}
	// root node value is always first in preorder
	var node = new(TreeNode)
	node.Val = preorder[0]

	// now we need to construct left and right child nodes
	// find position of node.Val in inorder array
	// left side will be a inorder for left branch, right - for right
	var inorderPos = slices.Index(inorder, node.Val)
	inorderLeft := inorder[:inorderPos]
	inorderRight := inorder[inorderPos+1:]
	// now construct preorder arrays for each child
	preorderLeft := preorder[1 : inorderPos+1]
	preorderRight := preorder[inorderPos+1:]

	node.Left = buildTree(preorderLeft, inorderLeft)
	node.Right = buildTree(preorderRight, inorderRight)
	return node
}
