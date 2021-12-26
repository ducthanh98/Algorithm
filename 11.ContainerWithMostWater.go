package main

import "fmt"


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	area := 0
	for l < r {
		area = max(area,min(height[l],height[r]) * (r - l))
		if height[l] < height[r]{
			l++
		} else {
			r--
		}
	}
	return area
}

func main()  {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}