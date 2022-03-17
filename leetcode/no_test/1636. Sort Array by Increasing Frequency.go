package main

import (
	"sort"
)

func frequencySort(nums []int) []int {
	mapValue := make(map[int]int, 0)
	for _, num := range nums {
		mapValue[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if mapValue[nums[i]] < mapValue[nums[j]] {
			return true
		}
		return false
	})
	return nums
}

func main() {
	frequencySort([]int{1, 1, 2, 2, 2, 3})
}
