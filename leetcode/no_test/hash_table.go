package main

const length = 99

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type HashTable struct {
	HashTable []*Node
}

func Constructor() HashTable {
	return HashTable{
		HashTable: make([]*Node, length),
	}
}

func (this *HashTable) Put(key int, value int) {
	idx := key % length
	if this.HashTable[idx] == nil {

		this.HashTable[idx] = &Node{
			Key:   key,
			Value: value,
			Next:  nil,
		}

	}

	node := this.HashTable[idx]
	for {
		if node.Key == key {
			node.Value = value
			return
		}

		if node.Next == nil {

			node.Next = &Node{
				Key:   key,
				Value: value,
				Next:  nil,
			}
			return

		} else {

			node = node.Next

		}
	}
}

func (this *HashTable) Get(key int) int {
	idx := key % length
	if this.HashTable[idx] == nil {
		return -1
	}
	node := this.HashTable[idx]
	for node != nil {
		if node.Key == key {
			return node.Value
		}
		node = node.Next
	}
	return -1

}

func (this *HashTable) Remove(key int) {
	idx := key % length
	if this.HashTable[idx] == nil {
		return
	}

	node := this.HashTable[idx]
	if node.Key == key {
		this.HashTable[idx] = node.Next
	}
	prev := node
	node = node.Next
	for node != nil {
		if node.Key == key {
			node = node.Next
			prev.Next = node
			break
		}
		prev = node
		node = node.Next
	}
}

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(100, 2)
	obj.Put(199, 3)
	obj.Get(1)
	obj.Remove(100)
	obj.Put(100, 3)

}
