package lc14

import (
	"strings"
)

type Trie struct {
	root *Node
}

type Node struct {
	// value string
	value string
	// childs [26]*Node
	childs map[int32]*Node
	isLast bool
}

func NewNode(val string) *Node {
	return &Node{
		value:  val,
		isLast: false,
		// childs: [26]*Node{},
		childs: make(map[int32]*Node, 26),
	}
}

func (t *Trie) CountChildren(node *Node) int {
	// count := 0
	// for _, child := range node.childs {
	// 	if child != nil {
	// 		count++
	// 	}
	// }
	// return count
	return len(node.childs)
}

func (t *Trie) Insert(value string) {
	lowered := strings.ToLower(strings.ReplaceAll(value, " ", ""))
	curr := t.root
	// for i := 0; i < len(runes); i++ {
	for _, r := range lowered {
		// pref := lowered[i] - 'a'
		// pref := runes[i]
		if curr.childs[r] == nil {
			// curr.childs[pref] = NewNode(string(lowered[i]))
			curr.childs[r] = NewNode(string(r))
		}
		curr = curr.childs[r]
	}
	curr.isLast = true
}

func (t *Trie) Search(word string) bool {
	lowered := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	curr := t.root
	// for i := 0; i < len(lowered); i++ {
	for _, r := range lowered {
		// index := lowered[i] - 'a'
		if curr == nil || curr.childs[r] == nil {
			return false
		}
		curr = curr.childs[r]
	}

	return curr.isLast
}

func (t *Trie) StartsWith(prefix string) bool {
	var lowered = strings.ToLower(strings.ReplaceAll(prefix, " ", ""))
	curr := t.root
	// for i := 0; i < len(wordLower); i++ {
	for _, r := range lowered {
		// index := wordLower[i] - 'a'
		if curr.childs[r] == nil {
			return false
		}
		curr = curr.childs[r]
	}
	return true
}
