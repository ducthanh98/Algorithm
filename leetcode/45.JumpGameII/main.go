package _5JumpGameII

import (
	"github.com/ducthanh98/Algorithm/leetcode/utils"
)

func jump(nums []int) int {
	var result, end, farthest int
	length := len(nums) - 1
	for i, num := range nums {
		farthest = utils.Max(farthest, num+i)
		if i < length && i == end {
			result++
			end = farthest
		}
	}
	return result
}
