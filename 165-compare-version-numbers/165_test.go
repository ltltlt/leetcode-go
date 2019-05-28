package compare_version_number_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseInt(s string, start int) (int, int) {
	var length int
	var value int

	for i := start; i < len(s); i++ {
		v := int(s[i]) - '0'
		if v > 9 || v < 0 {
			break
		}
		value *= 10
		value += v
		length++
	}

	return value, length
}

func compareVersion(version1 string, version2 string) int {
	for i, j := 0, 0; i < len(version1) || j < len(version2); {
		v1, len1 := parseInt(version1, i)
		v2, len2 := parseInt(version2, j)
		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return 1
		}
		i += len1 + 1
		j += len2 + 1
	}
	return 0
}

func TestCompareVersion(t *testing.T) {
	assert.Equal(t, -1, compareVersion("0.1", "1.1"))
	assert.Equal(t, 1, compareVersion("1.0.1", "1"))
	assert.Equal(t, -1, compareVersion("7.5.2.4", "7.5.3"))
	assert.Equal(t, 0, compareVersion("1.01", "1.001"))
	assert.Equal(t, 0, compareVersion("1.0", "1.0.0"))
}
