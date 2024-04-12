package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []string
		expected [][]string
	}{
		{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			[][]string{{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}}},
		{[]string{"root/a 1.txt(abcd) 2.txt(efsfgh)", "root/c 3.txt(abdfcd)", "root/c/d 4.txt(efggdfh)"},
			[][]string{{"root/c/3.txt"}, {"root/c/d/4.txt"}, {"root/a/1.txt"}, {"root/a/2.txt"}}},
	}

	for _, test := range tests {
		assert.Equal(test.expected, findDuplicate(test.input), fmt.Sprintf("Input %v ", test.input))
	}
}
