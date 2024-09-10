package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{beginWord: "a", endWord: "c", wordList: []string{"a", "b", "c"}, expected: 2},
	}

	for _, test := range tests {
		assert.Equal(test.expected, ladderLength(test.beginWord, test.endWord, test.wordList), "Incorrect")
	}
}
