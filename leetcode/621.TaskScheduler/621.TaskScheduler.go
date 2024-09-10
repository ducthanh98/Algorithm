package main

import (
	"container/heap"
)

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	if i > j {
		return true
	}
	return true
}

// Swap swaps the elements with indexes i and j.
func (m MaxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

// Push and Pop are used to append and remove the last element of the slice
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	count := make(map[byte]int)

	for _, task := range tasks {
		count[task]++
	}

	mh := &MaxHeap{}
	heap.Init(mh)

	for _, v := range count {
		heap.Push(mh, v)
	}

	res := 0

	for mh.Len() > 0 {
		remain := make([]int, 0)
		cycle := n + 1

		for cycle > 0 && mh.Len() > 0 {

			freq := (heap.Pop(mh)).(int)
			if freq > 1 {
				remain = append(remain, freq-1)
			}
			cycle--
			res++
		}

		for _, i := range remain {
			heap.Push(mh, i)
		}

		if mh.Len() == 0 {
			break
		}
		res += cycle

	}

	return res
}
