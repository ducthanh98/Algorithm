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
			//k    int
		}
		expected int
	}{
		{input: struct {
			nums []int
			//k    int
		}{nums: []int{1, 1, 0, 1}}, expected: 3},
		{input: struct {
			nums []int
			//k    int
		}{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, expected: 5},
		{input: struct {
			nums []int
			//k    int
		}{nums: []int{1, 1, 1}}, expected: 2},
	}

	for _, test := range tests {
		assert.Equal(test.expected, longestSubarray(test.input.nums), fmt.Sprintf("Input %v ", test.input))
	}
}
