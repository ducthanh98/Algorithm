package HouseRobber

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	result := 0
	cache := make(map[int]int)
	dp(nums, 0, length, &result, cache)
	return result
}

func dp(nums []int, index, length int, result *int, cache map[int]int) int {
	if index >= length {
		return 0
	} else if index == length-1 {
		return nums[index]
	}

	if _, ok := cache[index]; ok {
		return cache[index]
	}
	maxRob := 0

	for i := index; i < length; i++ {
		rob := nums[i] + dp(nums, i+2, length, result, cache)
		if rob > maxRob {
			maxRob = rob
		}
	}

	cache[index] = maxRob
	if maxRob > *result {
		*result = maxRob
	}
	return maxRob
}
