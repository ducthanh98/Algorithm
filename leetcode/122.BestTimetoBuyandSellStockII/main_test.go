package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		//{[]int{0, 1, -2, -3, -4}, 3},
		//{[]int{-1, -2, -3, 0, 1}, 2},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, test := range tests {
		assert.Equal(test.expected, maxProfit(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
