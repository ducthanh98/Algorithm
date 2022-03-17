package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdditiveNumber(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected bool
	}{
		{"112358", true},
		{"199100199", true},
		{"1023", false},
	}

	for _, test := range tests {
		assert.Equal(isAdditiveNumber(test.input), test.expected)
	}

}
