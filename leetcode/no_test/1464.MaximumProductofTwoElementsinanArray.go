package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 2 {
		return (nums[0] - 1) * (nums[1] -1)
	}
	var a,b int
	if nums[0] > nums[1] {
		a,b = nums[0],nums[1]
	} else {
		b,a = nums[0],nums[1]
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > a {
			b = a
			a = nums[i]
		} else if nums[i] > b {
			b = nums[i]
		}
	}
	return (a-1) * (b-1)
}

func main(){
	fmt.Println(maxProduct([]int{3,4,5,2}))
}