package main

//Implement a LFU/LRU Cache.
//This question is about basic data structure double linked list

type Node struct {
	Key int
	Value int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	Capacity int
	Cache map[int]*Node
	Head *Node
	Tail *Node
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Cache: make(map[int]*Node),
	}
}


func (this *LRUCache) Get(key int) int {
	if this.Cache[key] != nil {
		this.Remove(this.Cache[key])
		this.Add(this.Cache[key])
		return this.Cache[key].Value
	}
	return -1
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.Head
	if this.Head != nil {
		this.Head.Prev = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node != this.Head {
		node.Prev.Next = node.Next
	} else {
		this.Head = node.Next
	}
	if node != this.Tail {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
}


func (this *LRUCache) Put(key int, value int)  {
	node ,ok := this.Cache[key]
	if ok  {
		node.Value = value
		this.Remove(node)
		this.Add(node)
	} else{
		node = &Node{Key: key,Value: value}
		this.Cache[key] = node
		this.Add(node)
	}
	if len(this.Cache) > this.Capacity {
		delete(this.Cache,this.Tail.Key)
		this.Remove(this.Tail)
	}


}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */