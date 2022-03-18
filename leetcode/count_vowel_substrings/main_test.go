package count_vowel_substrings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountVowelSubstrings(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected int
	}{
		{"aeiouu", 2},
		{"unicornarihan", 0},
		{"cuaieuouac", 7},
	}

	for _, test := range tests {
		assert.Equal(countVowelSubstrings(test.input), test.expected, fmt.Sprintf("Input %v ", test.input))
	}
}
