package MinCostClimbingStairs

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	cache := make(map[int]int)
	return min(dp(cost, 0, length, cache), dp(cost, 1, length, cache))
}

func dp(cost []int, index, length int, cache map[int]int) int {
	if index >= length {
		return 0
	} else if index == length {
		return cost[index]
	}
	if _, ok := cache[index]; ok {
		return cache[index]
	}

	cache[index] = cost[index] + min(dp(cost, index+1, length, cache), dp(cost, index+2, length, cache))
	return cache[index]
}
