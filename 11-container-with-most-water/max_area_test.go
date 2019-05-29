package max_area_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// use two pointer, left and right
// if height[left] < height[right]
// move left forward(maxarea include left is [left, right], ignore all other area include left)
// else if height[left] > height[right]
// move right backward(maxarea include right is [left, right], same as above)
// else height[left] == height[right]
// move left forward and move right backward

// this solution finally cost time complexity to O(n), and space complexity to O(1)

// there are brute force solution has time complexity O(n^2)
// and I can think of a dynamic programming solution have worst time complexity O(n^2)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var maxArea int
	for left, right := 0, len(height)-1; left < right; {
		lh, rh := height[left], height[right]
		maxArea = max(maxArea, (right-left)*min(lh, rh))

		if lh >= rh {
			right--
		}
		if lh <= rh {
			left++
		}
	}
	return maxArea
}

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 2, maxArea([]int{2, 2}))
	assert.Equal(t, 0, maxArea([]int{1}))
	assert.Equal(t, 0, maxArea(nil))
}
