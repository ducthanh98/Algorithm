package FrequencyoftheMostFrequentElement

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	length := len(nums)
	i := 0
	j := 0
	sum := 0
	for j = 0; j < length; j++ {
		sum += nums[j]
		if (j-i+1)*nums[j]-sum > k {
			sum -= nums[i]
			i++
		}
	}
	return j - i
}

func tribonacci(n int) int {
	mapDp := make(map[int]int)
	return exec(n, mapDp)
}

func exec(n int, mapDp map[int]int) int {
	fmt.Println(mapDp)
	if n == 0 {
		return 0

	}

	if n == 1 || n == 2 {
		return 1
	}
	if _, ok := mapDp[n]; ok {
		return mapDp[n]
	}
	mapDp[n] = exec(n-1, mapDp) + exec(n-2, mapDp) + exec(n-3, mapDp)

	return mapDp[n]
}
