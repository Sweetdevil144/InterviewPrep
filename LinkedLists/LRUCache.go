package main

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	cap         int
	cache       map[int]*Node
	left, right *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		left:  &Node{},
		right: &Node{},
	}
	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *Node) {
	prev, next := this.right.prev, this.right
	prev.next = node
	next.prev = node
	node.next = next
	node.prev = prev
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		delete(this.cache, key)
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.insert(node)

	if len(this.cache) > this.cap {
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}
