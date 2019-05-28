package strstr_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Robin-Karp algorithm
// time complexity is O(n) and space complexity is O(1)

// check if two string have the same length equal
func equal(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// use int64 to store hash, but what if needle is too long?
// in this case, there is still a hash value for the string,
// this value may be less than 0, this won't hurt,
// since same string still have same hash, and string not same still usually don't have same hash
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	const base = 128
	var hash int64
	for _, r := range needle {
		hash = hash*base + int64(r)
	}
	l := len(needle)
	var hash2 int64
	for i := 0; i < l; i++ {
		hash2 = hash2*base + int64(haystack[i])
	}
	highest := int64(math.Pow(base, float64(l-1)))
	for i := l; i <= len(haystack); i++ {
		if hash2 == hash && equal(needle, haystack[i-l:i]) {
			return i - l
		}
		if i == len(haystack) {
			break
		}
		hash2 = (hash2-highest*int64(haystack[i-l]))*base + int64(haystack[i])
	}
	return -1
}

func TestStrStr(t *testing.T) {
	assert.Equal(t, 2, strStr("hello", "ll"))
	assert.Equal(t, -1, strStr("aaaaa", "bba"))
	assert.Equal(t, 0, strStr("a", "a"))
	assert.Equal(t, 26, strStr("bbcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
}
