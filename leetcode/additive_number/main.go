package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 2 {
		return false
	}
	length := len(num)
	for i := 1; i < length; i++ {

		for j := i + 1; j < length; j++ {
			first := num[0:i]
			second := num[i:j]
			num := num[j:len(num)]

			if isValid(first, second, num, j, length) {
				return true
			}
		}

	}

	return false
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("convert int error: %v", err)
	}
	return i
}

func isValid(num1, num2, num string, index int, length int) bool {
	if index == length {
		return true
	}
	if (len(num1) > 1 && strings.HasPrefix(num1, "0")) || (len(num2) > 1 && strings.HasPrefix(num2, "0")) {
		return false
	}
	result := strconv.Itoa(toInt(num1) + toInt(num2))
	if !strings.HasPrefix(num, result) {
		return false
	}
	index += len(result)
	num = num[len(result):len(num)]
	return isValid(num2, result, num, index, length)
}

func main() {
	fmt.Println(isAdditiveNumber("112358"))
}
