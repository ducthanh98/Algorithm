package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicCaculator(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		s        string
		expected int
	}{
		//{s: "1 + 1", expected: 2},
		//{s: " 2-1 + 2 ", expected: 3},
		{s: "(1+(4+5+2)-3)+(6+8)", expected: 23},
	}

	for _, test := range tests {
		assert.Equal(test.expected, calculate(test.s), fmt.Sprintf("Incorrect %v", test.s))
	}
}
