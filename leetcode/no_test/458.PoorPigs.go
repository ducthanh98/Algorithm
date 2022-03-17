package main

import (
	"fmt"
	"math"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	attempts := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(attempts))))
}

func main() {
	fmt.Println(poorPigs(8, 15, 40))
}
