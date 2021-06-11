package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	key := binarySearch(nums, 0, len(nums)-1, target)
	if key == -1 {
		return []int{-1, -1}
	}
	i, j := key, key
	for i >= 0 && nums[i] == target {
		i--
	}

	for j < len(nums) && nums[j] == target {
		j++
	}

	return []int{i + 1, j - 1}
}

func binarySearch(nums []int, start, end, target int) int {

	for start <= end {
		pivot := (start + end) / 2
		fmt.Println(start, end, pivot)

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] > target {
			return binarySearch(nums, start, pivot-1, target)
		} else {
			return binarySearch(nums, pivot+1, end, target)
		}
	}
	return -1
}

func main() {
	fmt.Println(searchRange([]int{2, 2}, 3))
}
