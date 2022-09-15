package _567_MaximumLengthofSubarrayWithPositiveProduct

import (
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	mapValue := make(map[string]int, 0)
	mapValue["I"] = 1
	mapValue["V"] = 5
	mapValue["X"] = 10
	mapValue["L"] = 50
	mapValue["C"] = 100
	mapValue["D"] = 500
	mapValue["M"] = 1000
	result := 0
	parsedInput := []rune(s)
	length := len(parsedInput)

	for i := 0; i < length; i++ {
		char := string(parsedInput[i])
		if i < length-1 {
			nextChar := string(parsedInput[i+1])
			if char == "I" && (nextChar == "V" || nextChar == "X") ||
				char == "X" && (nextChar == "L" || nextChar == "C") ||
				char == "C" && (nextChar == "D" || nextChar == "M") {
				i++
				result += mapValue[nextChar] - mapValue[char]
				continue
			}
		}
		result += mapValue[char]
	}
	return result

}
func Pow(a, b int64) int64 {
	var result int64 = 1
	var i int64
	for i = 0; i < b; i++ {
		result *= a
	}
	return result
}

func plusOne(digits []int) []int {
	var num int64 = 0
	length := int64(len(digits))
	for i, d := range digits {
		num += int64(d) * Pow(10, length-int64(i)-1)
	}
	num += 1
	str := strconv.FormatInt(num, 10)
	arr := strings.Split(str, "")
	result := make([]int, 0)
	for _, s := range arr {
		tmp, _ := strconv.Atoi(s)
		result = append(result, tmp)
	}
	return result
}
