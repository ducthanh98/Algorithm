package maximumProductSubarray

import "github.com/ducthanh98/Algorithm/leetcode/utils"

//  Largest Sum Contiguous Subarray:Kadane’s algorithm
func maxProduct(nums []int) int {
	maxEnding := nums[0]
	minEnding := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		temp := utils.Max(utils.Max(nums[i], nums[i]*maxEnding), nums[i]*minEnding)
		minEnding = utils.Min(utils.Min(nums[i], nums[i]*maxEnding), nums[i]*minEnding)
		maxEnding = temp
		maxSoFar = utils.Max(maxSoFar, maxEnding)
	}
	return maxSoFar
}
