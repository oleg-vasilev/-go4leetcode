package lc117

// 117. Populating Next Right Pointers in Each Node II

// Given a binary tree
//
// struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
// }
// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
//
// Initially, all next pointers are set to NULL.

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var queue = make([]*Node, 0)
	queue = append(queue, root)
	var levelItems int

	for len(queue) > 0 {
		levelItems = len(queue)
		for i := 0; i < levelItems; i++ {
			if i+1 < levelItems {
				queue[i].Next = queue[i+1]
			} else {
				queue[i].Next = nil
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levelItems:]
	}
	return root
}
