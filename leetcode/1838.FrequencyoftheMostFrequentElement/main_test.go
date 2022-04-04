package FrequencyoftheMostFrequentElement

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Input struct {
	Nums []int
	K    int
}

func TestFrequencyoftheMostFrequentElement(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    Input
		expected int
	}{
		{Input{Nums: []int{1, 2, 4}, K: 5}, 3},
		{Input{Nums: []int{1, 4, 8, 13}, K: 5}, 2},
		{Input{Nums: []int{3, 9, 6}, K: 2}, 1},
	}

	for _, test := range tests {
		assert.Equal(maxFrequency(test.input.Nums, test.input.K), test.expected, fmt.Sprintf("Input %v ", test.input))
	}
}
