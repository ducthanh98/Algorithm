package _567_MaximumLengthofSubarrayWithPositiveProduct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInit(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		assert.Equal(test.expected, romanToInt(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
