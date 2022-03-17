package main

import (
	"fmt"
	"math"
)

// func minSpeedOnTime(dist []int, hour float64) int {
// 	if float64(len(dist))-1 > hour {
// 		return -1
// 	}

// 	maxNumber := math.Pow(10, 7)
// 	i := float64(1)
// 	lastValue := -1
// 	for {
// 		result := float64(0)

// 		for j := 0; j < len(dist); j++ {
// 			if j == len(dist)-1 {
// 				result += float64(dist[j]) / float64(i)
// 				fmt.Println(result)
// 			} else {
// 				result += math.Ceil(float64(dist[j]) / float64(i))
// 			}

// 			if result > hour {
// 				i = math.Floor((maxNumber + i) / 2)
// 				break
// 			}
// 		}

// 		if result <= hour {
// 			i = math.Floor((1 + i) / 2)
// 		}

// 	}
// 	return -1
// }

var result = -1

func minSpeedOnTime(dist []int, hour float64) int {
	if float64(len(dist))-1 > hour {
		return -1
	}

	end := int(math.Pow(10, 7))
	start := 1

	findMinSpeed(dist, hour, start, end)

	return result
}

func findMinSpeed(dist []int, hour float64, start, end int) {
	if start > end {
		return
	}
	mid := (start + end) / 2
	var tmp float64 = 0
	flag := true
	for j := 0; j < len(dist); j++ {
		if j == len(dist)-1 {
			tmp += float64(dist[j]) / float64(mid)
		} else {
			tmp += math.Ceil(float64(dist[j]) / float64(mid))
		}

		if tmp > hour {
			if mid+1 <= end {
				findMinSpeed(dist, hour, mid+1, end)
			}
			flag = false
			break
		}
	}

	if flag {
		result = int(mid)
		findMinSpeed(dist, hour, start, mid-1)
	}

}

func main() {
	fmt.Println(minSpeedOnTime([]int{1, 1}, 1.0))

}
