package utils

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(*i)
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
