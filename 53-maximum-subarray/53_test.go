package maximum_subarray_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// dynamic programming solution
// which have time complexity O(n) and space complexity O(1)
// based on: Solution(nums[:i+1]) = max(Solution(nums[:i]), a suffix sum + nums[i])
// here's how to find this suffix:
// when we find new element, if previous suffix sum is 0, reset suffix to 0 and add
// this new element to suffix
func maxSubArray2(nums []int) int {
	var sum = math.MinInt64
	var suffixSum int
	for _, num := range nums {
		if suffixSum < 0 {
			suffixSum = 0
		}
		suffixSum += num
		sum = max(sum, suffixSum)
	}
	return sum
}

// divide and conquer
// divide nums into left part and right part
// now Solution(nums) = max(Solution(left), Solution(right), long sequence include boundary)
// no good solution to find the boundary-included sequence, may need O(n) to find it
// which left the time complexity to O(n logn); also, we have space complexity O(1)
// emmm, 4ms and 3.3MB but previous solution cost 8ms and 3.3MB, 3.3MB maybe golang runtime cost
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}
	m := l / 2
	left := nums[:m]
	right := nums[m:]
	return max(maxSubArray(left), max(
		maxSubArray(right),
		maxSuffix(left)+maxPrefix(right),
	))
}

func maxSuffix(nums []int) int {
	var suffixSum int
	for _, num := range nums {
		if suffixSum < 0 {
			suffixSum = 0
		}
		suffixSum += num
	}
	return suffixSum
}

func maxPrefix(nums []int) int {
	var prefixSum int
	for i := len(nums) - 1; i >= 0; i-- {
		if prefixSum < 0 {
			prefixSum = 0
		}
		prefixSum += nums[i]
	}
	return prefixSum
}

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, -1, maxSubArray([]int{-1}))
	assert.Equal(t, -1, maxSubArray([]int{-2, -1}))
}

func TestMaxSubArray2(t *testing.T) {
	assert.Equal(t, 6, maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, -1, maxSubArray2([]int{-1}))
	assert.Equal(t, -1, maxSubArray2([]int{-2, -1}))
}
