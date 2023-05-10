package main

func longestOnes(nums []int, k int) int {
	result, start, count := 0, 0, 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			count++
		}
		for count > k {
			if nums[start] == 0 {
				count--
			}
			start++
		}

		if result < (end - start + 1) {
			result = end - start + 1
		}
	}

	return result
}
