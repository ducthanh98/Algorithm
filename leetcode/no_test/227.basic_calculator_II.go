package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func calculate(s string) int {
	re := regexp.MustCompile(`(?m)\d+`)
	numbers := re.FindAllString(s, -1)

	re = regexp.MustCompile(`(?m)[+\-*/]`)
	operations := re.FindAllString(s, -1)
	stack := make([]int, 0)
	tmp, _ := strconv.Atoi(numbers[0])
	stack = append(stack, tmp)
	numbers = numbers[1:]
	stPosition := 0

	length := len(numbers)
	for i := 0; i < length; i++ {
		tmp, _ := strconv.Atoi(numbers[i])

		if operations[i] == "*" {
			stack[stPosition] = stack[stPosition] * tmp
		} else if operations[i] == "/" {
			stack[stPosition] = stack[stPosition] / tmp
		} else if operations[i] == "-" {
			stack = append(stack, -tmp)
			stPosition++
		} else {
			stack = append(stack, tmp)
			stPosition++
		}

	}

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func main() {
	fmt.Println(calculate("3/2"))
}
