package main
//Implement a LFU/LRU Cache.
//This question is about basic data structure double linked list


type LFUCache struct {
	Cache map[int]*Node
	Cap int
	Head *Node
	Tail *Node
}

type Node struct {
	Key int
	Val int
	Count int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LFUCache {
	c := LFUCache{
		Cache: make(map[int]*Node, capacity),
		Cap: capacity,
		Head: &Node{},
		Tail: &Node{},
	}
	c.Head.Next = c.Tail
	c.Tail.Prev = c.Head
	return c
}


func (this *LFUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}
	node.Count = node.Count + 1
	this.sort(node)
	return node.Val
}


func (this *LFUCache) Put(key int, value int) {
	if this.Cap == 0 {
		return
	}
	if this.isUpdate(key, value) {
		return
	}
	if len(this.Cache) > this.Cap {
		this.evict()
	}
	node := &Node{
		Key: key,
		Val: value,
		Count: 1,
	}
	this.Cache[key] = node
	this.toTail(node)
	this.sort(node)
}

func (this *LFUCache) isUpdate(key, value int) (b bool) {
	if node, ok := this.Cache[key]; ok {
		node.Val = value
		node.Count = node.Count + 1
		this.sort(node)
		return true
	}
	return
}

func (this *LFUCache) evict() {
	node := this.Tail.Prev

	node.Prev.Next = this.Tail
	this.Tail.Prev = node.Prev

	delete(this.Cache, node.Key)
}

func (this *LFUCache) toTail(node *Node) {
	prev := this.Tail.Prev

	prev.Next = node
	node.Prev = prev

	node.Next = this.Tail
	this.Tail.Prev = node
}

func (this *LFUCache) sort(node *Node) {
	index := node.Prev

	for index != this.Head && index.Count <= node.Count {
		index = index.Prev
	}
	if index.Next == node {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	next := index.Next
	index.Next = node
	node.Prev = index

	next.Prev = node
	node.Next = next
}