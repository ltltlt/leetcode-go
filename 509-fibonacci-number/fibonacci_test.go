package fibonacci_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// simple dynamic programming solution
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}

	solutions := make([]int, N+1)
	solutions[0] = 0
	solutions[1] = 1
	for i := 2; i <= N; i++ {
		solutions[i] = solutions[i-1] + solutions[i-2]
	}
	return solutions[N]
}

func TestFib(t *testing.T) {
	assert.Equal(t, 0, fib(0))
	assert.Equal(t, 1, fib(1))
	assert.Equal(t, 1, fib(2))
	assert.Equal(t, 2, fib(3))
	assert.Equal(t, 3, fib(4))
}
