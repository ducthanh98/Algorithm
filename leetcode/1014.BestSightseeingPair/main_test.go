package maximumProductSubarray

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{8, 1, 5, 2, 6}, 11},
		{[]int{1, 2}, 2},
	}

	for _, test := range tests {
		assert.Equal(test.expected, maxScoreSightseeingPair(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
