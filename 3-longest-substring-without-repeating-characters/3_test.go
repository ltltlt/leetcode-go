package longest_substring_test

// this problem is quite easy, simple dynamic programming

// solution(s+1) = when(s+1 already exist in previous suffix)
// solution(s)
// or
// max(solution(s), len(string's longest suffix)+1)

// O(n) time complexity and space complexity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) int {
	var length int
	var suffixLeft int
	set := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if pos, ok := set[b]; ok {
			for j := suffixLeft; j < pos+1; j++ {
				delete(set, s[j])
			}
			suffixLeft = pos + 1
		} else {
			length = max(length, i-suffixLeft+1)
		}
		set[b] = i
	}
	return length
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 5, lengthOfLongestSubstring("tmmzuxt"))
	assert.Equal(t, 6, lengthOfLongestSubstring("bbtablud"))
}
