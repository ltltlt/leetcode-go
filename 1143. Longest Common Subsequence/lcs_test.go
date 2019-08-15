package lcs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// simple dynamic programming problem
// since the input consist two string, let's build a table have size [text1.len+1][text2.len+1]
// each element in position i, j stores the longest common subsequence in text1[:i] and text2[:j]
// but how can we conclude table[i][j] from table[:i+1][:j+1] but not include this element
// case 1
// when text1[i] equals to text2[j], then table[i+1][j+1] should be table[i][j]+1 since another equal element
// included, since table[i][j] already is the lcs in text1[:i] and text2[:j], table[i][j]+1 must be the
// lcs in text1[:i+1] and text2[:j+1](if not, table[i+1][j] or table[i][j+1] must be larger than table[i][j]+1, this is impossible)
// case 2
// 当text1[i]!=text2[j]
// 当加入text1[i], 这个元素不和text2[j]相同， 所以这个元素只能和text2[0..j]的某个元素相配， 这个解包含在table[i+1][j]中, 或者是text2[j]和text1[0..i]某个元素相配，
// 这个解包含在table[i][j+1]中, 且相配不会在原先基础上+1， 因为没有一对新的相匹配

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	table := make([][]int, len(text1)+1)
	for i := 0; i < len(table); i++ {
		table[i] = make([]int, len(text2)+1)
	}

	for i := 0; i < len(text1); i++ {
		c := text1[i]
		for j := 0; j < len(text2); j++ {
			c2 := text2[j]
			if c == c2 {
				table[i+1][j+1] = table[i][j] + 1
			} else {
				table[i+1][j+1] = max(table[i][j+1], table[i+1][j])
			}
		}
	}

	return table[len(text1)][len(text2)]
}

// 现在table[i][j]存text1[:i]和text2[:j]的最长子字符串， 且这个子字符串包含text1[i-1]和text2[j-1]
// 最后需要遍历表得到一个最大值
func longestCommonSubstring(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	table := make([][]int, len(text1)+1)
	for i := 0; i < len(table); i++ {
		table[i] = make([]int, len(text2)+1)
	}

	for i := 0; i < len(text1); i++ {
		c := text1[i]
		for j := 0; j < len(text2); j++ {
			c2 := text2[j]
			if c == c2 {
				table[i+1][j+1] = table[i][j] + 1
			} else {
				table[i+1][j+1] = 0
			}
		}
	}

	var maximum int
	for _, array := range table {
		for _, elem := range array {
			maximum = max(maximum, elem)
		}
	}
	return maximum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestLCS(t *testing.T) {
	for i, c := range []struct {
		input1, input2 string
		output         int
		output2        int
	}{
		{"abcde", "ace", 3, 1},
		{"abc", "abc", 3, 3},
		{"abc", "def", 0, 0},
		{"", "", 0, 0},
		{"", "1", 0, 0},
	} {
		assert.Equalf(t, c.output, longestCommonSubsequence(c.input1, c.input2), "case %d fail", i)
		assert.Equalf(t, c.output2, longestCommonSubstring(c.input1, c.input2), "case %d fail 2", i)
	}
}
