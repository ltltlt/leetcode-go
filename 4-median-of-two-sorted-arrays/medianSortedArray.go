// pass, beat 100% golang submissions
// the code is simple, use two pointer for each array, record last two value
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		i, j    int
		average bool = (len(nums1)+len(nums2))%2 == 0
		middle  int  = (len(nums1) + len(nums2)) / 2
		v1, v2  int
	)
	for i+j <= middle {
		if j >= len(nums2) || (i < len(nums1) && nums1[i] < nums2[j]) {
			v1, v2 = v2, nums1[i]
			i++
		} else {
			v1, v2 = v2, nums2[j]
			j++
		}
	}
	if average {
		return float64(v1+v2) / 2
	} else {
		return float64(v2)
	}
}
