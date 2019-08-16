package roman_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = map[byte]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result int

	for i:=0; i<len(s); i++ {
		b := s[i]
		v := m[b]
		if i < len(s) -1 && isComposite(b, s[i+1]){
			v = m[s[i+1]] - v
			i++
		}
		result += v
	}

	return result
}

func isComposite(b1, b2 byte) bool {
	return (b1 == 'I' && (b2 == 'V' || b2 == 'X')) || 
	(b1 == 'X' && (b2 == 'L' || b2 == 'C')) || 
	(b1 == 'C' && (b2 == 'D' || b2 == 'M'))
}

func TestRomanToInt(t *testing.T) {
	for i, c := range []struct {
		input string
		output int
	}{
		{"III", 3},
		{"VI", 6},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}{
		assert.Equalf(t, c.output, romanToInt(c.input), "case %d fail", i)
	}
}