package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsage1(t *testing.T) {
	n1 := createListNode([]int{1, 8})
	n2 := createListNode([]int{0})
	addTwoNumbers(n1, n2)
}

func TestUsage2(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}

func TestUsage3(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m)
}

func TestUsage4(t *testing.T) {
	assert.Equal(t, "2314", getPermutation(4, 9))
}

func TestUsage5(t *testing.T) {
	nextPermutation([]int{3, 2, 1})
}
