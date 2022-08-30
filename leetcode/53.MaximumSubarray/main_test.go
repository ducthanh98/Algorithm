package _3_MaximumSubarray

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, test := range tests {
		assert.Equal(test.expected, maxSubArray(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
