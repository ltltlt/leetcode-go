package rematch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like . or *.

// we can use safely use byte array index for string

type Automata struct {
	s, p       string
	spos, ppos int
}

type State struct {
	b          byte // repeat byte
	spos, ppos int
}

type Byte byte // pattern byte

func (b Byte) match(b1 byte) bool {
	return b1 != 0 && (b == '.' || byte(b) == b1)
}

func NewAutomata(s, p string) *Automata {
	return &Automata{
		s: s,
		p: p,
	}
}

func (a *Automata) isMatch() bool {
	var asteriskStack []State
	for a.ppos < len(a.p) || a.spos < len(a.s) {
		pchar := getByte(a.p, a.ppos)
		// what happens if * is the first character in the p ?
		// or * appears after another * ?
		pnext := getByte(a.p, a.ppos+1)
		if pnext == '*' {
			a.ppos += 2
			asteriskStack = append(asteriskStack, State{
				b:    pchar,
				spos: a.spos,
				ppos: a.ppos,
			})
		} else if Byte(pchar).match(getByte(a.s, a.spos)) {
			a.spos++
			a.ppos++
		} else {
			// does not match
			// restore to previous state
			var state *State
			var b Byte
			for {
				l := len(asteriskStack)
				if l == 0 {
					return false
				}
				state = &asteriskStack[l-1]
				b = Byte(state.b)
				if b.match(getByte(a.s, state.spos)) {
					break
				}
				asteriskStack = asteriskStack[:l-1]
			}
			state.spos++
			a.ppos = state.ppos
			a.spos = state.spos
		}
	}

	if a.ppos == len(a.p) && a.spos == len(a.s) {
		return true
	}
	return false
}

func getByte(s string, i int) byte {
	if i < 0 || i >= len(s) {
		return 0
	}
	return s[i]
}

func isMatch(s string, p string) bool {
	a := NewAutomata(s, p)
	return a.isMatch()
}

func TestREMatch(t *testing.T) {
	assert.True(t, isMatch("aa", "aa"))
	assert.True(t, isMatch("aa", "a*"))
	assert.False(t, isMatch("ab", ".*c"))
}
