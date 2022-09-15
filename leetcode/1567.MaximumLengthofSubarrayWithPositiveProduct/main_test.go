package _567_MaximumLengthofSubarrayWithPositiveProduct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxLen(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []int
		expected int
	}{
		//{[]int{1, -2, -3, 4}, 4},
		//{[]int{0, 1, -2, -3, -4}, 3},
		//{[]int{-1, -2, -3, 0, 1}, 2},
		{[]int{5, -20, -20, -39, -5, 0, 0, 0, 36, -32, 0, -7, -10, -7, 21, 20, -12, -34, 26, 2}, 8},
	}

	for _, test := range tests {
		assert.Equal(test.expected, getMaxLen(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
