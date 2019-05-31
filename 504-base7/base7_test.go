package base7_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(bs []byte) []byte {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return bs
}

func convertToBase7(num int) string {
	var s []byte
	n := num
	if n < 0 {
		n = -n
	}
	for {
		s = append(s, byte('0'+n%7))
		n /= 7
		if n == 0 {
			break
		}
	}
	if num < 0 {
		s = append(s, '-')
	}
	return string(reverse(s))
}

func TestBase7(t *testing.T) {
	assert.Equal(t, "202", convertToBase7(100))
	assert.Equal(t, "-10", convertToBase7(-7))
	assert.Equal(t, "0", convertToBase7(0))
}
