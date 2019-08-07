package single_element_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func singleNonDuplicate(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

func TestSingle(t *testing.T) {
	for i, c := range []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 2, 2, 3, 3, 4, 5, 5}, 4},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 2},
		{nil, 0},
		{[]int{11}, 11},
	} {
		assert.Equalf(t, singleNonDuplicate(c.input), c.output, "case %v fail", i)
	}
}
