package _5JumpGameII

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJump(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}

	for _, test := range tests {
		assert.Equal(test.expected, jump(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
