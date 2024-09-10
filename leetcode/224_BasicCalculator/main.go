package main

import (
	"unicode"
)

func calculate(s string) int {
	var stacks []int32

	var res int32 = 0
	var signal int32 = 1
	var number int32 = 0

	for _, input := range s {

		if unicode.IsDigit(input) {
			number = 10*number + (input - '0')
		} else if input == '+' {
			res += signal * number
			number = 0
			signal = 1
		} else if input == '-' {
			res += signal * number
			number = 0
			signal = -1
		} else if input == '(' {
			stacks = append(stacks, res)
			stacks = append(stacks, signal)
			res = 0
			number = 0
			signal = 1
		} else if input == ')' {
			tmp := stacks[len(stacks)-2:]
			stacks = stacks[:len(stacks)-2]
			res += signal * number
			res = tmp[0] + tmp[1]*res
			signal = 1
			number = 0
		}

	}
	if number != 0 {
		res += signal * number
	}
	return int(res)

}
