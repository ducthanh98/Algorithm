package main

import (
	"container/heap"
)

// MaxHeap represents a max heap, inheriting heap.Interface
type MaxHeap []int

// Implementing the sort.Interface to define the MaxHeap behavior

// Len returns the number of elements in the heap
func (h MaxHeap) Len() int {
	return len(h)
}

// Less compares two elements, we reverse the logic for a max heap
func (h MaxHeap) Less(i, j int) bool {
	return h[i] < h[j] // Note: greater than, for a max heap
}

// Swap swaps two elements in the heap
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the last element from the heap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	Small *MaxHeap
	Large *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		Small: &MaxHeap{},
		Large: &MaxHeap{},
	}

}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.Small, -1*num)
	this.Small.Len()
	if this.Small.Len() > 0 && this.Large.Len() > 0 && -1(*this.Small)[0] > (*this.Large)[0] {
		tmp := heap.Pop(this.Small)
		heap.Push(this.Large, -1*tmp.(int))
	}

	if this.Small.Len() > this.Large.Len()+1 {
		tmp := heap.Pop(this.Small)
		heap.Push(this.Large, -1*tmp.(int))
	} else if this.Large.Len() > this.Small.Len()+1 {
		tmp := heap.Pop(this.Large)
		heap.Push(this.Small, -1*tmp.(int))
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.Small.Len() > this.Large.Len() {
		return float64((*this.Small)[0])
	} else if this.Large.Len() > this.Small.Len() {
		return float64((*this.Large)[0])
	} else {
		return float64((*this.Small)[0]+(*this.Large)[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
