package HouseRobber

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{[1,1]}, 1},
	}

	for _, test := range tests {
		assert.Equal(rob(test.input), test.expected, fmt.Sprintf("Input %v ", test.input))
	}
}
