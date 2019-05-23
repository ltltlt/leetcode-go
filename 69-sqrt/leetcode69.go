package main

import "fmt"

// Newton's method
func mySqrt(x int) int {
	var next, prev float64
	var eps = 0.001
	next = 1.1
	for diff := next - prev; !(diff > -eps && diff < eps); diff = next - prev {
		prev = next
		next = next/2 + float64(x)/2/next
	}
	return int(next)
}
