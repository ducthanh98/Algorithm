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
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1}, 1},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}, 4173},
	}

	for _, test := range tests {
		assert.Equal(rob(test.input), test.expected, fmt.Sprintf("Input %v ", test.input))
	}
}
