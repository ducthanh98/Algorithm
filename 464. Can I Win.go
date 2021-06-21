package main

import (
	"fmt"
	"strings"
)

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal < 1 {
		return false
	}
	visited := make([]string, maxChoosableInteger)
	for i := 0; i < len(visited); i++ {
		visited[i] = "0"
	}
	sum := (1 + maxChoosableInteger) * maxChoosableInteger / 2
	if sum < desiredTotal {
		return false
	}
	mapValue := make(map[string]bool, 0)
	return dfs(desiredTotal, visited, mapValue)

}

func dfs(total int, visited []string, mapValue map[string]bool) bool {
	if total <= 0 {
		return false
	}
	key := strings.Join(visited, "")
	if _, ok := mapValue[key]; ok {
		return mapValue[key]
	}

	for i := 0; i < len(visited); i++ {
		if visited[i] == "0" {
			visited[i] = "1"
			if total < i+1 || !dfs(total-i-1, visited, mapValue) {
				mapValue[key] = true
				visited[i] = "0"
				return true
			}
			visited[i] = "0"
		}
	}
	mapValue[key] = false
	return false
}

func main() {
	fmt.Println(canIWin(10, 11))
}
