package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input struct {
			nums []int
			k    int
		}
		expected int
	}{
		{input: struct {
			nums []int
			k    int
		}{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2}, expected: 6},
		{input: struct {
			nums []int
			k    int
		}{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3}, expected: 10},
	}

	for _, test := range tests {
		assert.Equal(longestOnes(test.input.nums, test.input.k), test.expected, fmt.Sprintf("Input %v ", test.input))
	}
}
