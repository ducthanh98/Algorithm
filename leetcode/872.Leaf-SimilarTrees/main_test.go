package main

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
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
	}

	for _, test := range tests {
		assert.Equal(test.expected, maxProduct(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
