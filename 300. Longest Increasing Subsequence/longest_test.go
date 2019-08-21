package longest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// O(n^2) solution, dynamic programming
// s[i] stores longest subsequence in nums[0...i]
// ls[i] stores longest subsequence in nums[0...i] which includes nums[i]
// then s[i] = max(ls[i], s[i-1])
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	s := make([]int, len(nums))
	ls := make([]int, len(nums))

	s[0], ls[0] = 1, 1

	for i := 1; i < len(nums); i++ {
		ls[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				ls[i] = max(ls[i], ls[j]+1)
			}
		}
		s[i] = max(s[i-1], ls[i])
	}

	return s[len(s)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// see https://leetcode.com/problems/longest-increasing-subsequence/discuss/74824/JavaPython-Binary-search-O(nlogn)-time-with-explanation
// 用反证法可以证明tails是递增的, 如果非递增, 则存在i使tails[i] >= tails[i+1]
// 即长度为i+1的递增子序列最后一个元素会比长度为i的递增子序列小, 那么我们就可以用这个长度为i+1的子序列的前i个构成的子序列, 此子序列长度为i且最后一个元素值比tails[i]小, 则tails[i]不是长度为i的递增子序列中最后一个最小的那个, 矛盾
func improved(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	tails := make([]int, len(nums)+1)
	maxSize := 0
	for _, num := range nums {
		// bisect
		l, r := 0, maxSize // position r is not included
		for l < r {        // 最后必然l==r, 只要一开始l<=r; 找不到invariant...
			m := (l + r) / 2
			if tails[m] >= num {
				r = m
			} else {
				l = m + 1
			}
		}
		tails[l] = num
		if l+1 > maxSize {
			maxSize = l + 1
		}
	}
	return maxSize
}

func TestLIS(t *testing.T) {
	for i, c := range []struct {
		input  []int
		output int
	}{
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
		{nil, 0},
		{[]int{11, 12, 13, 14, 15, 6, 7, 8, 101, 18}, 6},
		{[]int{2, 2}, 1},
	} {
		assert.Equalf(t, c.output, lengthOfLIS(c.input), "case %d fail", i)
		assert.Equalf(t, c.output, improved(c.input), "case %d fail", i)
	}
}
