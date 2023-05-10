package main

func longestSubarray(nums []int) int {
	i, previousDeletedPos := 0, 0
	res := 0
	isDeleted := false

	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 && !isDeleted {
			previousDeletedPos = j
			isDeleted = true
		} else if nums[j] == 0 {
			i = previousDeletedPos + 1
			previousDeletedPos = j

		}
		if res < (j - i) {
			res = j - i
		}
	}
	return res
}
