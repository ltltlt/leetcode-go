package main

import "fmt"

func lengthOfLastWord(s string) int {
	var i = len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	length := i + 1
	for ; i >= 0 && s[i] != ' '; i-- {
	}
	return length - i - 1
}
