package valid_palindrome_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}
		for l < r && !isAlphaNum(s[r]) {
			r--
		}
		if l < r && !caseInsensitive(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(b byte) bool {
	r := rune(b)
	return unicode.IsUpper(r) || unicode.IsLower(r) || unicode.IsNumber(r)
}

func caseInsensitive(b1, b2 byte) bool {
	return unicode.ToUpper(rune(b1)) == unicode.ToUpper(rune(b2))
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("A man, a plan, a canal: Panama"))
	assert.False(t, isPalindrome("race a car"))
}
