package lp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// see https://segmentfault.com/a/1190000003914228

// O(n) time and O(n) space

// runtime beats 100% go submission

// 冗余换来清晰

func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}

	s1 := make([]byte, len(s)*2+1)

	for i := 0; i < len(s); i++ {
		s1[i*2+1] = s[i]
	}

	// now our palindrome can have one situation: middle is some char from s

	maxRight := -1
	var pos int                // the middle position that palindrome contains maxRight pos
	RL := make([]int, len(s1)) // RL[i]表示以i为中心的最长回文的半径

	for i := 0; i < len(s1); i++ {
		if i >= maxRight {
			_, r := extend(s1, i, i)
			maxRight = r
			pos = i
			RL[i] = r - i + 1
		} else {
			// i被达到maxRight的回文包含, 通过其对pos(这个回文的中心)的镜像, 我们可以减少计算LR[i]的开销
			// find mirror
			j := pos - (i - pos) // j surely >= 0
			rlj := RL[j]
			left := pos - (maxRight - pos) // left boundary of pos's palindrome string
			jleft := j - rlj + 1
			if jleft > left { // 小回文且不到边界
				RL[i] = rlj
			} else {
				_, r := extend(s1, i-(maxRight-i), maxRight)
				RL[i] = r - i + 1
				if r > maxRight {
					maxRight = r
					pos = i
				}
			}
		}
	}

	max := 0
	maxPos := -1
	for i, elem := range RL {
		if elem > max {
			max = elem
			maxPos = i
		}
	}
	max = max - 1
	maxPos /= 2
	radius := max / 2
	if radius*2 != max {
		return s[maxPos-radius : maxPos+radius+1]
	}
	return s[maxPos-radius : maxPos+radius]
}

// extend, assume s[i..j] is palindrome
func extend(s []byte, i, j int) (int, int) {
	i--
	j++
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - 1
}

func TestManacher(t *testing.T) {
	for i, c := range []struct {
		input  string
		output string
	}{
		{"ccc", "ccc"},
		{"aba", "aba"},
		{"abac", "aba"},
		{"abbac", "abba"},
		{"abcba", "abcba"},
		{"abba", "abba"},
		{"", ""},
		{"1", "1"},
	} {
		assert.Equalf(t, c.output, longestPalindrome3(c.input), "case %v fail", i)
	}
}
