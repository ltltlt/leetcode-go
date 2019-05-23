package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findNthDigit(n int) int {
	var lowBound, i int
	for {
		p := math.Pow(10.0, float64(i))
		highBound := lowBound + (i+1)*9*int(p)
		if highBound >= n {
			break
		}
		i, lowBound = i+1, highBound
	}
	offset := n - lowBound - 1
	v := offset/(i+1) + int(math.Pow(10.0, float64(i)))
	extra := (i + 1) - (offset % (i + 1)) - 1
	return (v / int(math.Pow(10.0, float64(extra)))) % 10
}

func TestFindNthDigit(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(findNthDigit(100), 5)
	assert.Equal(findNthDigit(9), 9)
	assert.Equal(findNthDigit(10), 1)
	assert.Equal(findNthDigit(11), 0)
	assert.Equal(findNthDigit(1000000000), 1)
}
