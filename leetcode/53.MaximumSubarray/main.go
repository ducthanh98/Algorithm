package _3_MaximumSubarray

func maxSubArray(nums []int) int {
	result := 0
	dp := make([]int, len(nums))

	for i, num := range nums {
		if i == 0 {
			dp[i] = num
			result = num
			continue
		}
		total := dp[i-1] + num
		if total > num {
			dp[i] = total
		} else {
			dp[i] = num
		}
		if result < dp[i] {
			result = dp[i]
		}

	}

	return result
}
