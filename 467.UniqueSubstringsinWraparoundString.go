package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findSubstringInWraproundString(p string) int {
	count := 0
	if len(p) == 0 {
		return 0
	}
	pre := 1
	maxLen := make([]int, 26)
	maxLen[p[0]-'a'] = 1

	for i := 1; i < len(p); i++ {
		if p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25 {
			pre++
		} else {
			pre = 1
		}
		maxLen[p[i]-'a'] = max(maxLen[p[i]-'a'], pre)
	}
	for _, i := range maxLen {
		count += i
	}
	return count

}

func main() {
	fmt.Println(findSubstringInWraproundString("cac"))
}
