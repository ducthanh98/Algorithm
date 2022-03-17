package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	count := 0
	backtracking(&count, nums, target, 0, 0)
	return count
}

func backtracking(count *int, nums []int, target, index, sum int) {
	if index == len(nums) {
		if sum == target {
			*count++
		}
		return
	}

	current := nums[index]
	sum += current
	backtracking(count, nums, target, index+1, sum)

	sum -= 2 * current
	backtracking(count, nums, target, index+1, sum)

}

func main() {
	fmt.Println(findTargetSumWays([]int{1}, 1))
}
