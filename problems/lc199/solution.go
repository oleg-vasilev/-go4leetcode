package lc199

// 199. Binary Tree Right Side View

// Given the root of a binary tree, imagine yourself standing
// on the right side of it, return the values of the nodes
// you can see ordered from top to bottom.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result = make([]int, 0)
	var bfs func(node *TreeNode, level int)
	bfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(result) < level {
			result = append(result, node.Val)
		}

		bfs(node.Right, level+1)
		bfs(node.Left, level+1)
	}
	bfs(root, 1)
	return result
}

func rightSideViewIterative(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		result = append(result, queue[l-1].Val)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
