package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		s        string
		t        string
		expected string
	}{
		{s: "ADOBECODEBANC", t: "ABC", expected: "BANC"},
	}

	for _, test := range tests {
		assert.Equal(test.expected, minWindow(test.s, test.t), fmt.Sprintf("Input %v ", test.s))
	}
}
