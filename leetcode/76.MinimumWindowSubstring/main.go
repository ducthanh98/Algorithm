package main

import (
	"math"
	"strings"
)

func minWindow(s string, t string) string {
	chars := strings.Split(t, "")
	parents := strings.Split(s, "")
	hashMap := make(map[string]int)
	for _, char := range chars {
		hashMap[char]++
	}

	start, end, counter, min, from := 0, 0, len(chars), math.MaxInt, 0

	for end < len(s) {

		if hashMap[parents[end]] > 0 {
			counter--
		}
		hashMap[parents[end]]--
		end++

		for counter == 0 {
			if end-start < min {
				min = end - start
				from = start
			}
			hashMap[parents[start]]++
			if hashMap[parents[start]] > 0 {
				counter++
			}
			start++
		}

	}
	if min == math.MaxInt {
		return ""
	} else {
		return strings.Join(parents[from:from+min], "")
	}
}
