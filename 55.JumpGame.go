package main

import "fmt"


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


func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i <= reachable && i < len(nums); i++ {
		reachable = max(reachable,i + nums[i])
	}
	return reachable >= len(nums) - 1
}

func main()  {
	fmt.Println(canJump([]int{2,3,1,1,4}))
}
