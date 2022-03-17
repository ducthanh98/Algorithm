package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	mapLetter := make(map[string][]string, 0)
	mapLetter["2"] = []string{"a", "b", "c"}
	mapLetter["3"] = []string{"d", "e", "f"}
	mapLetter["4"] = []string{"g", "h", "i"}
	mapLetter["5"] = []string{"j", "k", "l"}
	mapLetter["6"] = []string{"m", "n", "o"}
	mapLetter["7"] = []string{"p", "q", "r", "s"}
	mapLetter["8"] = []string{"t", "u", "v"}
	mapLetter["9"] = []string{"x", "w", "y", "z"}

	results := make([]string, 0)
	result := make([]string, 0)
	if digits == "" {
		return []string{}
	}
	strArr := strings.Split(digits, "")
	return backtracking(results, result, strArr, 0, mapLetter)
}

func backtracking(results []string, result []string, digits []string, index int, mapLetter map[string][]string) []string {
	if index == len(digits) {
		results = append(results, strings.Join(result, ""))
		return results
	}
	letters := mapLetter[digits[index]]
	for i := 0; i < len(letters); i++ {
		result = append(result, letters[i])
		results = backtracking(results, result, digits, index+1, mapLetter)
		result = result[:len(result)-1]
	}
	return results
}

func main() {
	fmt.Println(letterCombinations("23"))
}
