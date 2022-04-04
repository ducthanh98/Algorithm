package ClimbingStairs

func climbStairs(n int) int {
	cache := make(map[int]int, 0)
	dp(n, 0, cache)
	return cache[0]
}

func dp(n int, current int, cache map[int]int) int {
	temp := 0
	if _, ok := cache[current]; ok {
		return cache[current]
	}
	for i := 1; i < 3; i++ {
		current += i

		if current == n {
			temp++
		} else if current < n {
			temp += dp(n, current, cache)
		}
		current -= i
	}
	cache[current] = temp
	return temp
}
