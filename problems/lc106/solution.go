package lc106

import (
	"slices"
)

// 106. Construct Binary Tree from Inorder and Postorder Traversal

// Given two integer arrays inorder and postorder where inorder is the
// inorder traversal of a binary tree and postorder is the postorder
// traversal of the same tree, construct and return the binary tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// L,R - left and right subtree value(s), N - node value
	// inorder = L,N,R
	// preorder = N,L,R
	// postorder = L,R,N

	// base case
	if len(postorder) == 0 {
		return nil
	}
	// root node value is always last in postorder
	var node = new(TreeNode)
	node.Val = postorder[len(postorder)-1]
	// now we need to construct left and right child nodes
	// find position of node.Val in inorder array
	var inorderPos = slices.Index(inorder, node.Val)
	// left side will be a inorder for left branch, right - for right
	inorderLeft := inorder[:inorderPos]
	inorderRight := inorder[inorderPos+1:]
	// now figure out postorder arrays for each childs using node.Val position in inorder array
	// because inorder pos is basically size of left subtree
	postorderLeft := postorder[:inorderPos]                    // left subthee
	postorderRight := postorder[inorderPos : len(postorder)-1] // right subtree excludes right edge because it is root

	node.Left = buildTree(inorderLeft, postorderLeft)
	node.Right = buildTree(inorderRight, postorderRight)
	return node
}
