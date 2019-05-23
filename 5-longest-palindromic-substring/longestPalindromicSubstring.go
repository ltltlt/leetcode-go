// Beats 10% golang submissions
// Too long, a little complicate
// Note that this code is not strictness,
// i don't consider rune which contains more than 1 byte

// example: "aba"
// start from ab, ba and aba, then expand
package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var (
		newLen2derive map[[2]int]struct{}
		len2derive    = make(map[[2]int]struct{})
		newLen3derive map[[2]int]struct{}
		len3derive    = make(map[[2]int]struct{})
		i             int
	)
	for ; i < len(s)-2; i++ {
		j, k := i+1, i+2
		if s[i] == s[j] {
			len2derive[[2]int{i, j}] = struct{}{}
		}
		if s[i] == s[k] {
			len3derive[[2]int{i, k}] = struct{}{}
		}
	}
	if s[i] == s[i+1] {
		len2derive[[2]int{i, i + 1}] = struct{}{}
	}
	for {
		newLen2derive = make(map[[2]int]struct{})
		for indexs, _ := range len2derive {
			if indexs[0] > 0 && indexs[1] < len(s)-1 && s[indexs[0]-1] == s[indexs[1]+1] {
				newLen2derive[[2]int{indexs[0] - 1, indexs[1] + 1}] = struct{}{}
			}
		}
		if len(newLen2derive) == 0 {
			break
		}
		len2derive = newLen2derive
	}
	for {
		newLen3derive = make(map[[2]int]struct{})
		for indexs, _ := range len3derive {
			if indexs[0] > 0 && indexs[1] < len(s)-1 && s[indexs[0]-1] == s[indexs[1]+1] {
				newLen3derive[[2]int{indexs[0] - 1, indexs[1] + 1}] = struct{}{}
			}
		}
		if len(newLen3derive) == 0 {
			break
		}
		len3derive = newLen3derive
	}
	var (
		indexs1, indexs2 *[2]int
	)
	if len(len3derive) != 0 {
		// random select key
		for indexs, _ := range len3derive {
			indexs1 = &indexs
			break
		}
	}
	if len(len2derive) != 0 {
		// random select key
		for indexs, _ := range len2derive {
			indexs2 = &indexs
			break
		}
	}
	if indexs1 != nil && indexs2 != nil {
		if indexs1[1]-indexs1[0] > indexs2[1]-indexs2[0] {
			return s[indexs1[0] : indexs1[1]+1]
		} else {
			return s[indexs2[0] : indexs2[1]+1]
		}
	} else if indexs1 != nil {
		return s[indexs1[0] : indexs1[1]+1]
	} else if indexs2 != nil {
		return s[indexs2[0] : indexs2[1]+1]
	} else {
		return string(s[0])
	}
}
