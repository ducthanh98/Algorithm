package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	used := make(map[int]bool, len(nums))
	results := make([][]int, 0)
	list := make([]int, 0)
	sort.Ints(nums)
	dfs(nums, used, &results, list)
	return results

}

func dfs(nums []int, used map[int]bool, results *[][]int, list []int) {
	if len(list) == len(nums) {
		var tmp = make([]int, len(nums))
		copy(tmp, list)
		*results = append(*results, tmp)
		return
	}

	for i, _ := range nums {
		if used[i] {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		used[i] = true
		list = append(list, nums[i])
		dfs(nums, used, results, list)
		fmt.Println("Used", used, "List", list, "results", results)
		used[i] = false
		list = list[:len(list)-1]
	}
}

func main() {
	fmt.Println(permuteUnique([]int{2, 2, 1, 1}))
}
