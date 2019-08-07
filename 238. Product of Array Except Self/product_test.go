package product_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// use two auxiliary array, array1[i] stores nums[i]'s left elements product,
// similarly, array2[i] stores nums[i]'s right elements product

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	left := make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(left); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i, elem := range right {
		left[i] *= elem
	}

	return left
}

func TestProduct(t *testing.T) {
	for i, c := range []struct {
		input  []int
		output []int
	}{
		{input: []int{2, 3, 4, 5}, output: []int{60, 40, 30, 24}},
		{input: nil, output: nil},
		{input: []int{2}, output: []int{1}},
	} {
		input := c.input
		output := c.output
		real := productExceptSelf(input)

		assert.ElementsMatchf(t, output, real, "case %v fail", i)
	}
}
