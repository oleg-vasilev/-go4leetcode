package lc129

// 129. Sum Root to Leaf Numbers

// You are given the root of a binary tree containing digits from 0 to 9 only.
//
// Each root-to-leaf path in the tree represents a number.
//
// For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
// Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.
//
// A leaf node is a node with no children.
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// easy recursive solution

// func sumNumbers(root *TreeNode) int {
// 	var dfs func(*TreeNode, int) int
//
// 	dfs = func(node *TreeNode, sum int) int {
// 		// base case
// 		if node == nil {
// 			return 0
// 		}
// 		// apply value to sum
// 		sum = sum*10 + node.Val
// 		// leaf reached
// 		if node.Left == nil && node.Right == nil {
// 			return sum
// 		}
// 		// go to left and right nodes
// 		return dfs(node.Left, sum) + dfs(node.Right, sum)
// 	}
//
// 	return dfs(root, 0)
// }

// iterative dfs

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var answer int

	var sumStack = make([]int, 0)
	var nodesStack = make([]*TreeNode, 0)

	nodesStack = append(nodesStack, root)
	sumStack = append(sumStack, 0)

	for len(nodesStack) > 0 {
		// pop top node
		item := nodesStack[len(nodesStack)-1]
		nodesStack = nodesStack[:len(nodesStack)-1]
		// pop top sum
		sum := sumStack[len(sumStack)-1]
		sumStack = sumStack[:len(sumStack)-1]
		// update sum with node value
		sum = sum*10 + item.Val
		// leaf node - add sum to answer and continue
		if item.Left == nil && item.Right == nil {
			answer += sum
			continue
		}
		// not a leaf - add childs to nodesStack with current sum
		if item.Right != nil {
			nodesStack = append(nodesStack, item.Right)
			sumStack = append(sumStack, sum)
		}
		if item.Left != nil {
			nodesStack = append(nodesStack, item.Left)
			sumStack = append(sumStack, sum)
		}
	}
	return answer
}
