// Beats 10% submissions.
// Too long and complicate.

// Since the problem tells us this string only contains
// digits, so we don't need to consider rune which contains
// more than 1 byte. And we just think rune as byte
package main

import (
	"fmt"
	"strings"
)

var (
	itor map[int]byte // convert int to byte
	rtoi map[byte]int // convert byte to int
)

func reverse(str string) string {
	res := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}
	return string(res)
}
func mulEach(num1, num2 string) []string {
	// return numbers in reverse order, eg: number 123 will be "321"
	var nums []string
	num1, num2 = reverse(num1), reverse(num2)
	for i, r := range num2 {
		stack := []byte{}
		carry := 0
		for _, rr := range num1 {
			x := rtoi[byte(rr)]*rtoi[byte(r)] + carry
			carry = x / 10
			stack = append(stack, itor[x%10])
		}
		if carry != 0 {
			stack = append(stack, itor[carry])
		}
		nums = append(nums, strings.Repeat("0", i)+string(stack))
	}
	return nums
}
func sumUp(nums []string) string {
	var (
		stack  []byte
		carry  int
		maxLen int
	)
	for _, num := range nums {
		if len(num) > maxLen {
			maxLen = len(num)
		}
	}
	for i := 0; i < maxLen; i++ {
		res := carry
		for _, num := range nums {
			if i < len(num) {
				res += rtoi[num[i]]
			}
		}
		carry = res / 10
		stack = append(stack, itor[res%10])
	}
	if carry != 0 {
		stack = append(stack, itor[carry])
	}
	return reverse(string(stack))
}
func multiply(num1 string, num2 string) string {
	itor = map[int]byte{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	}
	rtoi = map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	numbers := mulEach(num1, num2)
	return sumUp(numbers)
}
