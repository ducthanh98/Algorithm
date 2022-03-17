package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[j][0] > intervals[i][0]
	})

	newInterval := intervals[0]
	result := make([][]int, 0)
	result = append(result, newInterval)

	for _, interval := range intervals {
		if interval[0] <= newInterval[1] {
			newInterval[1] = max(newInterval[1], interval[1])
		} else {
			newInterval = interval
			result = append(result, newInterval)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge([][]int{[]int{1, 4}, []int{0, 2}, []int{3, 5}}))
}
