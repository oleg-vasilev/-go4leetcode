package lc146

// 146. LRU Cache

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
//
// Implement the LRUCache class:
//
// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.

type Entry struct {
	Key  int
	Val  int
	Prev *Entry
	Next *Entry
}

type LRUCache struct {
	data map[int]*Entry
	root *Entry
	len  int
	cap  int
}

func Constructor(capacity int) LRUCache {
	root := new(Entry) // dummy root node, head will be root.Next, last will be root.Prev
	root.Next = root
	root.Prev = root
	return LRUCache{
		data: make(map[int]*Entry),
		root: root,
		len:  0,
		cap:  capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	node, exists := c.data[key]
	if exists {
		// key exists - move requested key up in list
		moveAfter(node, c.root)
		return node.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// change and move existing entry
	node, exists := c.data[key]
	if exists {
		// key present - replace value in existed node
		node.Val = value
		// and move node to the front
		moveAfter(node, c.root)
		return
	}

	// key not present - create new node
	newNode := &Entry{Key: key, Val: value}
	c.data[key] = newNode
	insertAfter(newNode, c.root)
	c.len++
	// check if we should evict last item
	if c.len > c.cap {
		lastNode := c.root.Prev
		delete(c.data, lastNode.Key) // remove from data map
		remove(lastNode)             // remove from list
		c.len--
	}
}

func insertAfter(n, after *Entry) *Entry {
	n.Prev = after
	n.Next = after.Next
	n.Prev.Next = n
	n.Next.Prev = n
	return n
}

func moveAfter(n, after *Entry) {
	if n == after {
		return
	}
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Prev = after
	n.Next = after.Next
	n.Prev.Next = n
	n.Next.Prev = n
}

func remove(n *Entry) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Next = nil // avoid memory leaks
	n.Prev = nil // avoid memory leaks
}
