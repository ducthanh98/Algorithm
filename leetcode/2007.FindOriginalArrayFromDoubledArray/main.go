package main

import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	mapUsed := make(map[int]int, 0)

	results := make([]int, 0)
	for _, i := range changed {
		mapUsed[i]++
	}

	for _, i := range changed {
		val, _ := mapUsed[i]
		doubledValue, doubledExist := mapUsed[i*2]
		if val > 0 && !doubledExist && doubledValue < 1 {
			return []int{}
		} else if doubledExist && val > 0 && doubledValue > 0 {
			mapUsed[i]--
			mapUsed[i*2]--
			results = append(results, i)
		}
	}
	for _, k := range mapUsed {
		if k != 0 {
			return []int{}
		}
	}

	return results
}

func main() {
	fmt.Println(findOriginalArray([]int{0, 0, 0, 0, 0, 0}))
}
