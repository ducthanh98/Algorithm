package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type Median struct {
	Head *Node
	Count int
	Total int
}

func (this *Median) addNum(num int)  {
	node := &Node{
		Value: num,
		Next:  this.Head,
	}
	this.Count++
	this.Total = this.Total + num

	this.Head = node
}

func (this *Median) findMedian() float64 {
	return float64(this.Total)/float64(this.Count)
}

func main()  {
	median := &Median{
		Head:  nil,
		Count: 0,
		Total: 0,
	}

	median.addNum(1)
	median.addNum(2)
	median.findMedian()
	median.addNum(3)

	median.findMedian()
	node := median.Head
	for node != nil {
		fmt.Print(node.Value)
		node = node.Next
	}

}
