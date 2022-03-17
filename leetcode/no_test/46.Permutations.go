package main

import "fmt"

var visited map[int]bool
var results [][]int

func backtracking(nums []int, result []int) {

	if len(nums) == len(result) {
		tmp := make([]int, len(nums))
		copy(tmp, result)
		results = append(results, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			result = append(result, nums[i])
			backtracking(nums, result)
			result = result[:len(result)-1]
			visited[i] = false
		}

	}
}

func permute(nums []int) [][]int {
	visited = make(map[int]bool, len(nums))
	results = make([][]int, 0)

	backtracking(nums, make([]int, 0))
	return results
}

func main() {
	fmt.Println(permute([]int{0, 1, 2}))

}
