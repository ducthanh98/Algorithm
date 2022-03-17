package main

import "fmt"

func findPeakElement(nums []int) int {
	max := nums[0]
	idx := 0
	for i := 1; i < len(nums)-2; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] && nums[i] > max {
			idx = i
			max = nums[i]
		}
	}
	if nums[len(nums)-1] > max {
		return len(nums) - 1
	}
	return idx
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1}))
}
