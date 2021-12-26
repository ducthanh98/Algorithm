package main

import (
	"fmt"
	"sort"
)

var results = make([][]int,0)
var temp = make([]int,0)
func combinationSum(candidates []int, target int) [][]int {
	results = make([][]int,0)
	temp = make([]int,0)

	sort.Ints(candidates)

	backtrack(candidates,target,0)
	return results
}

func backtrack(candidates []int, target,index int) {

	for i := index; i < len(candidates); i++ {
		remain := target - candidates[i]
		if remain > 0 {
			temp = append(temp,candidates[i])
			backtrack(candidates,remain,i)
			temp = temp[:len(temp)-1]
		} else if remain == 0 {
			temp = append(temp,candidates[i])
			dist := make([]int,len(temp))
			copy(dist,temp)
			results = append(results,dist)
			temp = temp[:len(temp)-1]
			break
		} else {
			break
		}
	}
}

func main()  {
	fmt.Println(combinationSum([]int{2,3,5},8))
}
