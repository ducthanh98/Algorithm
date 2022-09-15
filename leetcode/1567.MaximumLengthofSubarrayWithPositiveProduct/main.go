package _567_MaximumLengthofSubarrayWithPositiveProduct

import (
	"github.com/ducthanh98/Algorithm/leetcode/utils"
)

func getMaxLen(nums []int) int {
	result := 0
	negative := 0
	isArrayContainNegativeValue := false
	firstNegative := -1
	positive := 0

	for i, num := range nums {
		if num == 0 {
			negative = 0
			positive = 0
			isArrayContainNegativeValue = false
			firstNegative = -1

		} else if num > 0 {
			positive++
			negative++
		} else if isArrayContainNegativeValue {
			negative++
			positive = negative
			isArrayContainNegativeValue = false
		} else if !isArrayContainNegativeValue {
			negative++
			positive = 0
			isArrayContainNegativeValue = true
			if firstNegative == -1 {
				firstNegative = i
			}
		}
		result = utils.Max(result, positive)
		if isArrayContainNegativeValue {
			result = utils.Max(result, i-firstNegative)
		}
	}
	return result
}
