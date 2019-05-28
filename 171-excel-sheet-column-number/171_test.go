package column_number_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func titleToNumber(s string) int {
	var num int
	weight := 1
	for i := len(s) - 1; i >= 0; i-- {
		b := s[i]
		num += weight * (int(b) - 'A' + 1)
		weight *= 26
	}
	return num
}

func TestTitleToNumber(t *testing.T) {
	assert.Equal(t, 26, titleToNumber("Z"))
	assert.Equal(t, 27, titleToNumber("AA"))
}
