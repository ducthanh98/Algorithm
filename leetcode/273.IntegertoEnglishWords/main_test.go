package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberToWord(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected string
	}{
		{123, "One Hundred Twenty Three"},
		{12345, "Twelve Thousand Three Hundred Forty Five"},
		{20, "Twenty"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	}

	for _, test := range tests {
		assert.Equal(test.expected, numberToWords(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
