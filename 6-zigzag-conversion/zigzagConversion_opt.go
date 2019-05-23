package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	var (
		buffer bytes.Buffer
		step   [2]int
	)
	for i := 0; i < numRows; i++ {
		step[0], step[1] = 2*(numRows-i)-2, 2*i
		if step[0] == 0 {
			if step[1] == 0 { // when numRows equal to 1
				step = [...]int{1, 1}
			} else {
				step[0] = step[1]
			}
		} else if step[1] == 0 {
			step[1] = step[0]
		}
		for j, k := i, true; ; k = !k {
			if j < len(s) {
				buffer.WriteByte(s[j])
			} else {
				break
			}
			if k {
				j += step[0]
			} else {
				j += step[1]
			}
		}
	}
	return buffer.String()
}
