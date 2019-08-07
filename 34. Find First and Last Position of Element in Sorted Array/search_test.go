package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchRange(nums []int, target int) []int {
	pos := bisect(nums, target)
	if pos == -1 {
		return []int{-1, -1}
	}
	var left, right int
	for left = pos; left >= 0 && nums[left] == target; left-- {
	}
	for right = pos; right < len(nums) && nums[right] == target; right++ {
	}
	return []int{left + 1, right - 1}
}

func bisect(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func TestSearchRange(t *testing.T) {
	for i, c := range []struct {
		input1 []int
		input2 int
		output []int
	}{
		{[]int{1, 2, 3, 3, 4, 5}, 3, []int{2, 3}},
		{[]int{1, 2, 3, 3, 4, 5}, -1, nil},
		{nil, -1, nil},
	} {
		assert.ElementsMatchf(t, c.output, searchRange(c.input1, c.input2), "case %v fail", i)
	}
}
