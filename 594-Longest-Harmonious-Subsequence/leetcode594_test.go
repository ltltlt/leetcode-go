package leetcode594

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

type pattern struct {
	i, j int
	one  bool
}

func (p pattern) match(i int) bool {
	if p.one {
		return abs(i-p.i) < 2
	} else {
		return i == p.i || i == p.j
	}
}

func (p pattern) add(j int) pattern {
	if j != p.i {
		p.j = j
		p.one = false
	}
	return p
}

func findLHS(nums []int) int {
	var m = make(map[pattern]int, len(nums))
	for _, num := range nums {
		var patterns = make(map[pattern]int)
		for pattern, c := range m {
			if pattern.match(num) {
				p := pattern.add(num)
				if c1, ok := patterns[p]; ok {
					c = max(c, c1)
				}
				patterns[p] = c
			}
		}
		for pattern, c := range patterns {
			m[pattern] = c + 1
		}
		p := pattern{i: num, one: true}
		if _, ok := m[p]; !ok {
			m[p] = 1
		}
	}
	var longest = 0
	for p, l := range m {
		if !p.one {
			longest = max(longest, l)
		}
	}
	return longest
}

func TestFindLHS(t *testing.T) {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	assert.Equal(t, 5, findLHS2(nums), "error")
}
