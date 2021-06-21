package main

import (
	"fmt"
	"math"
	"sort"
)

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	return []int{minStep(stones), maxStep(stones)}
}

func minStep(stones []int) int {
	left, right := 0, 1
	minStep := math.MaxInt32
	gaps := 0
	existStones := 1
	length := len(stones)
	for right < length {

		if stones[right] <= stones[left]+length-1 {
			right++
			existStones++
		} else {

			gaps = length - existStones
			if gaps == 1 && stones[right-1] == stones[left]+length-1 || gaps != 1 {
				minStep = min(gaps, minStep)
			}
			left++
			existStones--

		}

	}
	gaps = length - existStones
	if gaps == 1 && stones[right-1] == stones[left]+length-1 || gaps != 1 {
		minStep = min(gaps, minStep)
	}
	return minStep

}

func maxStep(stones []int) int {
	total := stones[len(stones)-1] - stones[0] + 1
	space := total - len(stones)
	return max(space-(stones[1]-stones[0]-1), space-(stones[len(stones)-1]-stones[len(stones)-2]-1))
}

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

func main() {
	fmt.Println(numMovesStonesII([]int{4, 7, 50}))
}
