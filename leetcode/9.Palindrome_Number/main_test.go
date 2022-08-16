package Palindrome_Number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{699, false},
	}

	for _, test := range tests {
		assert.Equal(isPalindrome(test.input), test.expected, fmt.Sprintf("Input %v ", test.input))
	}
}
