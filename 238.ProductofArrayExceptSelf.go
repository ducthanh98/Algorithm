package main

func productExceptSelf(nums []int) []int {
	multiplyResult := 1
	countZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			multiplyResult *= nums[i]
		} else {
			countZero++
		}
	}

	results := make([]int, len(nums))
	if countZero > 1 {
		return results
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			results[i] = multiplyResult
		} else if countZero == 1 {
			results[i] = 0
		} else {
			results[i] = multiplyResult / nums[i]
		}
	}

	return results

}
