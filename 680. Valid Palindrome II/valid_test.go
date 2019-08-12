package valid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validPalindrome(s string) bool {
	return validPalindrome2(s, 0, len(s)-1, true)
}

func validPalindrome2(s string, start, end int, canDelete bool) bool {
	if end <= start {
		return true
	}
	if s[start] == s[end] {
		return validPalindrome2(s, start+1, end-1, canDelete)
	}

	if !canDelete {
		return false
	}

	return validPalindrome2(s, start+1, end, false) || validPalindrome2(s, start, end-1, false)
}

func TestValidPalindrome(t *testing.T) {
	for i, c := range []struct {
		input  string
		output bool
	}{
		{"abca", true},
		{"abc", false},
		{"", true},
		{"1", true},
		{"12", true},
		{"aba", true},
	} {
		assert.Equalf(t, c.output, validPalindrome(c.input), "case %d fail", i)
	}
}
