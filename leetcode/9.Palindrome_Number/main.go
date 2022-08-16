package Palindrome_Number

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	end := len(str) - 1
	start := 0
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--

	}
	return true
}
